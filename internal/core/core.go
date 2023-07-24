package core

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/shurco/litecart/pkg/jwtutil"
)

// Flags is ...
type Flags struct {
	Serve bool
}

// Core is ...
type Core struct {
	DevMode    bool
	MainDomain string
	Log        zerolog.Logger
	Fiber      *fiber.App
	Stripe     *Stripe
	JWT        *jwtutil.Setting
}
