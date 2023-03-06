package httpfx

import (
	"github.com/gofiber/fiber/v2"
	"github.com/veteran-dev/veteran/pkg/system/configfx"
	"go.uber.org/fx"
)

func Setup(c *configfx.Config) *fiber.App {
	return fiber.New(fiber.Config{
		// Prefork:       true,
		// CaseSensitive: true,
		// StrictRouting: true,
		// ServerHeader:  "Fiber",
		AppName: c.AppName,
	})
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(Setup),
)
