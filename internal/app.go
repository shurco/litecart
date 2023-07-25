package app

import (
	"embed"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/armon/go-proxyproto"
	"github.com/rs/zerolog"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/internal/routes"
	"github.com/shurco/litecart/pkg/fsutil"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

var (
	DevMode    bool
	MainDomain string
)

type Flags struct {
	Serve bool
}

// NewApp is ...
func NewApp(flags Flags) error {
	DevMode = true
	log := zerolog.New(os.Stderr).With().Timestamp().Logger()

	if flags.Serve {
		if err := fsutil.MkDirs(0775, "./uploads"); err != nil {
			log.Err(err).Send()
			return err
		}

		if err := queries.InitQueries(embedMigrations); err != nil {
			log.Err(err).Send()
			return err
		}
		db := queries.DB()

		// web web server
		var views *html.Engine
		if DevMode {
			views = html.New("../web/views", ".html")
			views.Reload(true)
		} else {
			views = html.NewFileSystem(http.Dir("../web/views"), ".html")
		}
		views.Delims("{#", "#}")
		views.AddFunc(
			"arr", func(els ...any) []any {
				return els
			},
		)

		app := fiber.New(fiber.Config{
			DisableStartupMessage: true,
			Views:                 views,
		})

		app.Use(fiberzerolog.New(fiberzerolog.Config{
			Logger: &log,
		}))

		app.Static("/", "../web/public")
		app.Static("/uploads", "./uploads")

		app.Use(func(c *fiber.Ctx) error {
			// init install
			mainPath := strings.Split(c.Path(), "/")[1]
			if !db.IsInstalled() {
				if c.Path() != "/_/install" && mainPath != "api" {
					return c.Redirect("/_/install")
				}
			} else if c.Path() == "/_/install" {
				return c.Redirect("/_")
			}

			// init main domain
			if MainDomain == "" {
				hostname := strings.Split(c.Hostname(), ".")
				if len(hostname) > 2 {
					hostname = hostname[1:]
				}
				MainDomain = strings.Join(hostname, ".")
			}

			// check subdomain
			if len(c.Subdomains()) > 0 {
				if !db.CheckSubdomain(c.Subdomains()[0]) && !DevMode {
					return c.Redirect(fmt.Sprintf("%s://%s", c.Protocol(), MainDomain), fiber.StatusMovedPermanently)
				}
			}

			return c.Next()
		})

		routes.SiteRoutes(app)
		routes.AdminRoutes(app)
		routes.ApiRoutes(app)
		routes.NotFoundRoute(app)

		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
		if err != nil {
			log.Err(err).Send()
		}
		proxyListener := &proxyproto.Listener{Listener: ln}
		if err := app.Listener(proxyListener); err != nil {
			log.Err(err).Send()
		}
	}

	return nil
}
