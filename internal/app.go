package app

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/armon/go-proxyproto"
	"github.com/rs/zerolog"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/internal/routes"
	"github.com/shurco/litecart/migrations"
	"github.com/shurco/litecart/pkg/fsutil"
	"github.com/shurco/litecart/pkg/logging"
	"github.com/shurco/litecart/web"
)

var (
	DevMode    bool
	MainDomain string
	log        zerolog.Logger
)

// NewApp is ...
func NewApp(appDev bool) error {
	DevMode = appDev
	log = logging.Log()

	// check lc_uploads folder
	if err := fsutil.MkDirs(0775, "./lc_uploads"); err != nil {
		log.Err(err).Send()
		return err
	}

	if err := queries.InitQueries(migrations.Embed()); err != nil {
		log.Err(err).Send()
		return err
	}

	// web web server
	var views *html.Engine
	if DevMode {
		views = html.New("../web/site", ".html")
		views.Reload(true)
	} else {
		if fsutil.IsDir("./web") {
			views = html.New("./web", ".html")
			views.Reload(true)
		} else {
			views = html.NewFileSystem(http.FS(web.Embed()), ".html")
		}
	}
	views.Delims("{#", "#}")

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Views:                 views,
	})

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &log,
	}))

	app.Static("/", "../web/site/public")
	app.Static("/components", "../web/site/components")
	app.Static("/uploads", "./lc_uploads")

	app.Use(DatabaseCheck)
	app.Use(SubdomainCheck)

	routes.SiteRoutes(app)
	routes.AdminRoutes(app)
	routes.ApiPrivateRoutes(app)
	routes.ApiPublicRoutes(app)
	routes.NotFoundRoute(app)

	if DevMode {
		StartServer(app)
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

		StartServer(app)
		<-idleConnsClosed
	}

	return nil
}

func DatabaseCheck(c *fiber.Ctx) error {
	db := queries.DB()
	if !db.IsInstalled() {
		if !strings.HasPrefix(c.Path(), "/_/install") && !strings.HasPrefix(c.Path(), "/api") {
			return c.Redirect("/_/install")
		}
	} else if strings.HasPrefix(c.Path(), "/_/install") {
		return c.Redirect("/_")
	}
	return c.Next()
}

func SubdomainCheck(c *fiber.Ctx) error {
	db := queries.DB()

	/*
		if MainDomain == "" {
			hostname := strings.Split(c.Hostname(), ".")
			if len(hostname) > 2 {
				hostname = hostname[1:]
			}
			MainDomain = strings.Join(hostname, ".")
		}
	*/

	if len(c.Subdomains()) > 0 {
		if !db.CheckSubdomain(c.Subdomains()[0]) && !DevMode {
			return c.Redirect(fmt.Sprintf("%s://%s", c.Protocol(), MainDomain), fiber.StatusMovedPermanently)
		}
	}
	return c.Next()
}

// StartServer is ...
func StartServer(a *fiber.App) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Err(err).Send()
	}
	proxyListener := &proxyproto.Listener{Listener: ln}
	if err := a.Listener(proxyListener); err != nil {
		log.Err(err).Send()
	}
}
