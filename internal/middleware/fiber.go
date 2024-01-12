package middleware

import (
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
)

// FiberMiddleware is ...
func Fiber(a *fiber.App, log *zerolog.Logger) {
	a.Use(cors.New())
	a.Use(helmet.New())
	a.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	a.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: log,
	}))
	a.Use(recover.New())
}
