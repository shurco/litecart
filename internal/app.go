package app

import (
	"crypto/tls"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/rs/zerolog"
	"golang.org/x/crypto/acme/autocert"

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
	DevMode    bool
	MainDomain string
	log        zerolog.Logger
)

// NewApp is ...
func NewApp(httpAddr, httpsAddr string, appDev bool) error {
	DevMode = appDev
	log = logging.Log()

	schema := "http"
	mainAddr := httpAddr
	if httpsAddr != "" {
		schema = "https"
		mainAddr = httpsAddr
	}

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
		//Prefork:               true,
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

	fmt.Print("🛒 litecart - open source shopping-cart in 1 file\n")
	fmt.Printf("├─ Cart UI: %s://%s/\n", schema, mainAddr)
	fmt.Printf("└─ Admin UI: %s://%s/_/\n", schema, mainAddr)

	if schema == "https" {
		m := &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(mainAddr),
			Cache:      autocert.DirCache("./lc_certs"),
		}

		cfgTLS := &tls.Config{
			GetCertificate: m.GetCertificate,
			NextProtos: []string{
				"http/1.1", "acme-tls/1",
			},
		}

		ln, err := tls.Listen("tcp", ":443", cfgTLS)
		if err != nil {
			log.Err(err).Send()
			os.Exit(1)
		}

		if err := app.Listener(ln); err != nil {
			log.Err(err).Send()
			os.Exit(1)
		}

	}

	if DevMode {
		StartServer(mainAddr, app)
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

		StartServer(mainAddr, app)
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
func StartServer(addr string, a *fiber.App) {
	if err := a.Listen(addr); err != nil {
		log.Err(err).Send()
		os.Exit(1)
	}
}
