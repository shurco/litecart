package routes

import (
	"io/fs"
	"strings"

	"github.com/gofiber/fiber/v2"

	handlers "github.com/shurco/litecart/internal/handlers/public"
	"github.com/shurco/litecart/web"
)

// SiteRoutes sets up routes for the public-facing website SPA.
func SiteRoutes(c *fiber.App) {
	embedSite, _ := fs.Sub(web.EmbedSite(), web.SiteBuildPath)

	// Payment success and cancel routes: process payment first, then pass to SPA
	// These are NOT API endpoints - they handle redirects from payment providers
	// Use middleware approach to allow c.Next() to work properly
	c.Use("/cart/payment/success", handlers.PaymentSuccess)
	c.Use("/cart/payment/cancel", handlers.PaymentCancel)

	// Skip API routes, admin routes (/_ but not /_app), uploads
	// Note: /cart/payment/success and /cart/payment/cancel are NOT skipped
	// because they need to be handled by SPA after payment processing
	skipPaths := func(path string) bool {
		return strings.HasPrefix(path, "/api") ||
			(path == "/_" || (strings.HasPrefix(path, "/_/") && !strings.HasPrefix(path, "/_app"))) ||
			strings.HasPrefix(path, "/uploads")
	}

	c.Use("/", setupSPAHandler(embedSite, skipPaths))
}
