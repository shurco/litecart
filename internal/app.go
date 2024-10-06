package app

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/acme/autocert"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/shurco/litecart/internal/middleware"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/internal/routes"
	"github.com/shurco/litecart/migrations"
	"github.com/shurco/litecart/pkg/fsutil"
	"github.com/shurco/litecart/pkg/logging"
	"github.com/shurco/litecart/pkg/webutil"
	"github.com/shurco/litecart/web"
)

var (
	DevMode bool
	log     *logging.Log
)

// NewApp is ...
func NewApp(httpAddr, httpsAddr string, noSite, appDev bool) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	DevMode = appDev
	log = logging.New()

	schema := "http"
	mainAddr := httpAddr
	if httpsAddr != "" {
		schema = "https"
		mainAddr = httpsAddr
	}

	if err := queries.New(migrations.Embed()); err != nil {
		log.Err(err).Send()
		return err
	}

	// web web server
	fiberConfig := fiber.Config{
		// Prefork:               true,
		DisableStartupMessage: true,
		BodyLimit:             50 * 1024 * 1024,
	}
	var sitePath string

	if !noSite {
		sitePath = "./site"
		if DevMode {
			sitePath = "../web/site"
		} else {
			if !fsutil.IsDir(sitePath) || fsutil.IsEmptyDir(sitePath) {
				fsutil.EmbedExtract(web.EmbedSite(), "")
			}
		}
		views := html.New(sitePath, ".html")
		views.Reload(true)
		views.Delims("{#", "#}")
		fiberConfig.Views = views
	}

	app := fiber.New(fiberConfig)
	middleware.Fiber(app, log.Logger)

	// init structure
	if err := Init(); err != nil {
		log.Err(err).Send()
		os.Exit(1)
	}
	app.Static("/uploads", "./lc_uploads")
	app.Use(InstallCheck)
	routes.AdminRoutes(app)
	routes.ApiPrivateRoutes(app)

	fmt.Print("🛒 litecart - open source shopping-cart in 1 file\n")
	if !noSite {
		app.Static("/", sitePath+"/public", fiber.Static{
			CacheDuration: 30 * 24 * time.Hour,
		})
		routes.SiteRoutes(app)
		routes.ApiPublicRoutes(app)
		fmt.Printf("├─ Cart UI: %s://%s/\n", schema, mainAddr)
	}
	fmt.Printf("└─ Admin UI: %s://%s/_/\n", schema, mainAddr)

	routes.NotFoundRoute(app, noSite)

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
		if err := StartServer(ctx, mainAddr, app); err != nil {
			log.Err(err).Send()
		}
	} else {
		idleConnsClosed := make(chan struct{})

		go func() {
			sigint := make(chan os.Signal, 1)
			signal.Notify(sigint, os.Interrupt)
			<-sigint

			cancel()

			if err := app.Shutdown(); err != nil {
				log.Err(err).Send()
			}

			close(idleConnsClosed)
		}()

		go func() {
			if err := StartServer(ctx, mainAddr, app); err != nil {
				log.Err(err).Send()
			}
		}()

		<-idleConnsClosed
	}

	return nil
}

func InstallCheck(c *fiber.Ctx) error {
	db := queries.DB()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := db.GetSettingByKey(ctx, "installed")
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}
	install, _ := strconv.ParseBool(response["installed"].Value.(string))

	if !install {
		if !strings.HasPrefix(c.Path(), "/_/install") && !strings.HasPrefix(c.Path(), "/_/assets") && !strings.HasPrefix(c.Path(), "/api") {
			return c.Redirect("/_/install")
		}
	} else if strings.HasPrefix(c.Path(), "/_/install") {
		return c.Redirect("/_")
	}
	return c.Next()
}

// StartServer is ...
func StartServer(ctx context.Context, addr string, a *fiber.App) error {
	errCh := make(chan error)

	go func() {
		if err := a.Listen(addr); err != nil {
			log.Err(err).Send()
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		err := errors.New("shutdown signal received, closing server")
		log.Err(err).Send()
		return a.Shutdown()
	case err := <-errCh:
		return err
	}
}
