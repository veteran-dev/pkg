package bundlefx

import (
	"context"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/veteran-dev/pkg/system/configfx"
	"github.com/veteran-dev/pkg/system/dbfx"
	"github.com/veteran-dev/pkg/system/httpfx"
	"github.com/veteran-dev/pkg/system/loggerfx"
	"github.com/veteran-dev/pkg/system/redisfx"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func registerHooks(
	lifecycle fx.Lifecycle,
	slog *zap.SugaredLogger, cfg *configfx.Config, app *fiber.App, db *gorm.DB, redis *redis.Client,
) {
	idleConnsClosed := make(chan struct{})
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					sigint := make(chan os.Signal, 1)
					signal.Notify(sigint, os.Interrupt) // Catch OS signals.
					<-sigint
					slog.Info("Gracefully shutting down...")

					// Received an interrupt signal, shutdown.
					if err := app.Shutdown(); err != nil {
						// Error from closing listeners, or context timeout:
						slog.Errorf("Server is not shutting down! Reason: %v", err)
					}
					close(idleConnsClosed)
				}()

				go func() {
					slog.Infof("Server running, port=%s", cfg.ApplicationConfig.Address)
					if err := app.Listen(cfg.ApplicationConfig.Address); err != nil {
						slog.Errorf("Server is not running! Reason: %v", err)
					}
				}()

				return nil
			},
			OnStop: func(context.Context) error {
				redis.Close()
				<-idleConnsClosed
				return slog.Sync()
			},
		},
	)
}

// Module provided to fx
var Module = fx.Options(
	configfx.Module,
	loggerfx.Module,
	httpfx.Module,
	dbfx.Module,
	redisfx.Module,
	fx.Invoke(registerHooks),
)
