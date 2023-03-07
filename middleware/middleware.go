package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(),
		limiter.New(limiter.Config{
			Max:        100,
			Expiration: 1 * time.Second,
			KeyGenerator: func(c *fiber.Ctx) string {
				return c.IP()
			},
			LimitReached: func(c *fiber.Ctx) error {
				return c.SendStatus(fiber.StatusTooManyRequests)
			},
			SkipFailedRequests:     false,
			SkipSuccessfulRequests: false,
			LimiterMiddleware:      limiter.FixedWindow{},
		}),
		// Add simple logger.
		logger.New(logger.Config{
			TimeFormat: time.RFC3339,
			TimeZone:   "Asia/Shanghai",
			Done: func(c *fiber.Ctx, logString []byte) {
				if c.Response().StatusCode() != fiber.StatusOK {
					c.Send(logString)
				}
			},
		}),
		recover.New(),
	)
}
