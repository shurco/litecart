package app

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/armon/go-proxyproto"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/shurco/litecart/internal/app/routes"
)

const (
	webPort        = 8080
	templateExt    = ".html"
	templateViews  = "../web/views"
	templatePublic = "../web/public"
	uploadsPublic  = "./uploads"
)

func (b *Base) initWebServer(port int) error {
	var views *html.Engine
	if b.DevMode {
		views = html.New(templateViews, templateExt)
		views.Reload(true)
	} else {
		views = html.NewFileSystem(http.Dir(templateViews), templateExt)
	}
	views.Delims("{#", "#}")
	views.AddFunc(
		"arr", func(els ...any) []any {
			return els
		},
	)

	b.Fiber = fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Views:                 views,
	})

	b.Fiber.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &b.Log,
	}))

	b.Fiber.Static("/", templatePublic)
	b.Fiber.Static("/uploads", uploadsPublic)

	b.Fiber.Use(func(c *fiber.Ctx) error {
		// init install
		mainPath := strings.Split(c.Path(), "/")[1]
		if !b.DB.IsInstalled() {
			if c.Path() != "/_/install" && mainPath != "api" {
				return c.Redirect("/_/install")
			}
		} else if c.Path() == "/_/install" {
			return c.Redirect("/_")
		}

		// init main domain
		if b.MainDomain == "" {
			hostname := strings.Split(c.Hostname(), ".")
			if len(hostname) > 2 {
				hostname = hostname[1:]
			}
			b.MainDomain = strings.Join(hostname, ".")
		}

		// check subdomain
		if len(c.Subdomains()) > 0 {
			if !b.DB.CheckSubdomain(c.Subdomains()[0]) && !b.DevMode {
				return c.Redirect(fmt.Sprintf("%s://%s", c.Protocol(), b.MainDomain), fiber.StatusMovedPermanently)
			}
		}

		return c.Next()
	})

	routes.SiteRoutes(&b.Core, b.DB)
	routes.AdminRoutes(&b.Core, b.DB)
	routes.ApiRoutes(&b.Core, b.DB)
	routes.NotFoundRoute(b.Fiber)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", webPort))
	if err != nil {
		b.Log.Err(err).Send()
	}
	proxyListener := &proxyproto.Listener{Listener: ln}
	if err := b.Fiber.Listener(proxyListener); err != nil {
		b.Log.Err(err).Send()
	}

	return nil
}
