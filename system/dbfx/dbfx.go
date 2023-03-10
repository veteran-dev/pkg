package dbfx

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/veteran-dev/pkg/system/configfx"
	"go.uber.org/fx"
)

func ConnectRedis(c *configfx.Config) *redis.Client {
	addr := fmt.Sprintf("%s:%d", c.DatabaseConfig.Host, c.DatabaseConfig.Port)

	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: c.DatabaseConfig.Password, // no password set
		DB:       c.DatabaseConfig.DB,       // use default DB
	})
}

type Logger struct {
	Name string
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(ConnectRedis),
	fx.Provide(func() *Logger {
		return &Logger{Name: "redis"}
	}),
)

const (
	AccountToken = "account:token:%d"
)
