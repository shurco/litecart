package routes

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shurco/litecart/web"
)

// AdminRoutes sets up routes for the admin panel SPA.
func AdminRoutes(c *fiber.App) {
	embedAdmin, _ := fs.Sub(web.EmbedAdmin(), "admin/build")

	// Custom handler to properly serve static files and fallback to index.html for SPA routes
	c.Use("/_", func(c *fiber.Ctx) error {
		path := strings.TrimPrefix(c.Path(), "/_")
		if path == "" || path == "/" {
			path = "index.html"
		} else {
			// Remove leading slash for fs.Open
			path = strings.TrimPrefix(path, "/")
		}

		// Try to open the file from embed FS
		file, err := embedAdmin.Open(path)
		if err == nil {
			defer file.Close()

			// Check if it's a directory
			stat, err := file.Stat()
			if err == nil && stat.IsDir() {
				// For directories, serve index.html
				file.Close()
				file, err = embedAdmin.Open("index.html")
				if err != nil {
					return c.Next()
				}
				defer file.Close()
				path = "index.html"
			}

			// Determine content type based on file extension
			ext := filepath.Ext(path)
			contentType := getContentType(ext)
			c.Set("Content-Type", contentType)

			// Serve the file
			return c.SendStream(file)
		}

		// If file not found and it's a static asset request, return 404
		// Otherwise, serve index.html for SPA routing
		if isStaticAsset(path) {
			return c.Status(fiber.StatusNotFound).SendString("Not Found")
		}

		// For SPA routes, serve index.html
		indexFile, err := embedAdmin.Open("index.html")
		if err != nil {
			return c.Next()
		}
		defer indexFile.Close()
		c.Set("Content-Type", "text/html")
		return c.SendStream(indexFile)
	})
}

// isStaticAsset checks if the path is a static asset (JS, CSS, images, etc.).
func isStaticAsset(path string) bool {
	staticExts := []string{".js", ".css", ".png", ".jpg", ".jpeg", ".gif", ".svg", ".ico", ".woff", ".woff2", ".ttf", ".eot", ".json"}
	ext := strings.ToLower(filepath.Ext(path))
	for _, staticExt := range staticExts {
		if ext == staticExt {
			return true
		}
	}
	return false
}

// getContentType returns the appropriate content type for a file extension.
func getContentType(ext string) string {
	ext = strings.ToLower(ext)
	contentTypes := map[string]string{
		".html":  "text/html",
		".js":    "application/javascript",
		".css":   "text/css",
		".json":  "application/json",
		".png":   "image/png",
		".jpg":   "image/jpeg",
		".jpeg":  "image/jpeg",
		".gif":   "image/gif",
		".svg":   "image/svg+xml",
		".ico":   "image/x-icon",
		".woff":  "font/woff",
		".woff2": "font/woff2",
		".ttf":   "font/ttf",
		".eot":   "application/vnd.ms-fontobject",
	}

	if ct, ok := contentTypes[ext]; ok {
		return ct
	}
	return "application/octet-stream"
}
