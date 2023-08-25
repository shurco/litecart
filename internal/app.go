package app

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"

	"github.com/pires/go-proxyproto"
	"github.com/rs/zerolog"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/template/html/v2"

	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/internal/routes"
	"github.com/shurco/litecart/migrations"
	"github.com/shurco/litecart/pkg/fsutil"
	"github.com/shurco/litecart/pkg/logging"
	"github.com/shurco/litecart/web"
)

var (
	ProxyMod   bool
	DevMode    bool
	MainDomain string
	log        zerolog.Logger
)

// NewApp is ...
func NewApp(proxyMod, appDev bool) error {
	ProxyMod = proxyMod
	DevMode = appDev
	log = logging.Log()

	if err := queries.InitQueries(migrations.Embed()); err != nil {
		log.Err(err).Send()
		return err
	}

	// web web server
	var views *html.Engine
	var sitePath string

	if DevMode {
		sitePath = "../web/site"
		views = html.New(sitePath, ".html")
	} else {
		sitePath = "./site"
		if !fsutil.IsDir(sitePath) {
			fsutil.EmbedExtract(web.EmbedSite(), "")
		}
		views = html.New(sitePath, ".html")
	}
	views.Reload(true)
	views.Delims("{#", "#}")

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Views:                 views,
	})

	app.Use(helmet.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &log,
	}))

	app.Static("/", sitePath+"/public")
	app.Static("/components", sitePath+"/components")

	// check lc_uploads folder
	if err := fsutil.MkDirs(0775, "./lc_uploads"); err != nil {
		log.Err(err).Send()
		return err
	}
	app.Static("/uploads", "./lc_uploads")

	app.Use(InstallCheck)
	app.Use(SubdomainCheck)

	routes.AdminRoutes(app)
	routes.ApiPrivateRoutes(app)
	routes.SiteRoutes(app)
	routes.ApiPublicRoutes(app)

	routes.NotFoundRoute(app)

	if DevMode {
		StartServer(app, ProxyMod)
	} else {
		idleConnsClosed := make(chan struct{})

		go func() {
			sigint := make(chan os.Signal, 1)
			signal.Notify(sigint, os.Interrupt)
			<-sigint

			if err := app.Shutdown(); err != nil {
				log.Err(err).Send()
			}

			close(idleConnsClosed)
		}()

		StartServer(app, ProxyMod)
		<-idleConnsClosed
	}

	return nil
}

func InstallCheck(c *fiber.Ctx) error {
	db := queries.DB()
	if !db.IsInstalled() {
		if !strings.HasPrefix(c.Path(), "/_/install") && !strings.HasPrefix(c.Path(), "/_/assets") && !strings.HasPrefix(c.Path(), "/api") {
			return c.Redirect("/_/install")
		}
	} else if strings.HasPrefix(c.Path(), "/_/install") {
		return c.Redirect("/_")
	}
	return c.Next()
}

func SubdomainCheck(c *fiber.Ctx) error {
	//db := queries.DB()

	/*
		if MainDomain == "" {
			hostname := strings.Split(c.Hostname(), ".")
			if len(hostname) > 2 {
				hostname = hostname[1:]
			}
			MainDomain = strings.Join(hostname, ".")
		}
	*/

	/*
		if len(c.Subdomains()) > 0 {
			if !db.CheckSubdomain(c.Subdomains()[0]) && !DevMode {
				return c.Redirect(fmt.Sprintf("%s://%s", c.Protocol(), MainDomain), fiber.StatusMovedPermanently)
			}
		}
	*/
	return c.Next()
}

// StartServer is ...
func StartServer(a *fiber.App, proxyMode bool) {
	if proxyMode {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
		if err != nil {
			log.Err(err).Send()
		}

		proxyListener := &proxyproto.Listener{Listener: ln}
		defer proxyListener.Close()

		if err := a.Listener(proxyListener); err != nil {
			log.Err(err).Send()
		}
	} else {
		if err := a.Listen(fmt.Sprintf(":%d", 8080)); err != nil {
			log.Err(err).Send()
		}
	}

}
