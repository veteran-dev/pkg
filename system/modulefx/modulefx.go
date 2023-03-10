package modulefx

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/veteran-dev/pkg/system/configfx"
	"github.com/veteran-dev/pkg/system/loggerfx"
	"github.com/veteran-dev/pkg/system/sqlxfx"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func registerHooks(
	lifecycle fx.Lifecycle,
	slog *zap.SugaredLogger, cfg *configfx.Config, db *sqlx.DB,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {

				return nil
			},
			OnStop: func(context.Context) error {
				return slog.Sync()
			},
		},
	)
}

// Module provided to fx
var Module = fx.Options(
	configfx.Module,
	loggerfx.Module,
	// httpfx.Module,
	sqlxfx.Module,
	// redisfx.Module,
	fx.Invoke(registerHooks),
)
