package app

import (
	"embed"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/armon/go-proxyproto"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/internal/routes"
	"github.com/shurco/litecart/pkg/fsutil"
	"github.com/shurco/litecart/pkg/logging"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

var (
	DevMode    bool
	MainDomain string
)

// NewApp is ...
func NewApp() error {
	DevMode = true
	log := logging.Log()

	if err := fsutil.MkDirs(0775, "./uploads"); err != nil {
		log.Err(err).Send()
		return err
	}

	if err := queries.InitQueries(embedMigrations); err != nil {
		log.Err(err).Send()
		return err
	}

	// web web server
	var views *html.Engine
	if DevMode {
		views = html.New("../web/views", ".html")
		views.Reload(true)
	} else {
		views = html.NewFileSystem(http.Dir("../web/views"), ".html")
	}
	views.Delims("{#", "#}")

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Views:                 views,
	})

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &log,
	}))

	app.Static("/", "../web/public")
	app.Static("/uploads", "./uploads")

	app.Static("/_/components", "../web/views/admin/components")

	app.Use(DatabaseCheck)
	app.Use(SubdomainCheck)

	routes.SiteRoutes(app)
	routes.AdminRoutes(app)
	routes.ApiPrivateRoutes(app)
	routes.ApiPublicRoutes(app)
	routes.NotFoundRoute(app)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Err(err).Send()
	}
	proxyListener := &proxyproto.Listener{Listener: ln}
	if err := app.Listener(proxyListener); err != nil {
		log.Err(err).Send()
	}

	return nil
}

/*
func DatabaseCheck(c *fiber.Ctx) error {
	db := queries.DB()
	if !db.IsInstalled() {
		if c.Path() != "/_/install" && strings.Split(c.Path(), "/")[1] != "api" {
			return c.Redirect("/_/install")
		}
	} else if c.Path() == "/_/install" {
		return c.Redirect("/_")
	}
	return c.Next()
}
*/

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
