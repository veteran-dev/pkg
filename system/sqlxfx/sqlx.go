package sqlxfx

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/veteran-dev/pkg/system/configfx"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ConnectDB(c *configfx.Config, slog *zap.SugaredLogger) *sqlx.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", c.PostgresHost, c.PostgresUser, c.PostgresPassword, c.PostgresDbName, c.PostgresPort, c.PostgresSslMode, c.PostgresTimeZone)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		slog.Errorf("atabase connection failed, err=%s", err)
	}
	slog.Info("successfully connected to the database.")
	return db
}

var Module = fx.Options(
	fx.Provide(ConnectDB),
)
