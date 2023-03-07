package redisfx

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/veteran-dev/pkg/system/configfx"
	"go.uber.org/fx"
)

func ConnectRedis(c *configfx.Config, db int) *redis.Client {
	addr := fmt.Sprintf("%s:%d", c.RedisConfig.RedisHost, c.RedisConfig.RedisPort)
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: c.RedisConfig.RedisPassword, // no password set
		DB:       db,
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
