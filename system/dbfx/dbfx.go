package dbfx

import (
	"fmt"

	"github.com/albrow/zoom"
	"github.com/veteran-dev/pkg/system/configfx"
	"go.uber.org/fx"
)

func ConnectRedis(c *configfx.Config) *zoom.Pool {
	addr := fmt.Sprintf("%s:%d", c.DatabaseConfig.Host, c.DatabaseConfig.Port)

	options := zoom.PoolOptions{
		Address:  addr,
		Password: c.DatabaseConfig.Password,
		Database: c.DatabaseConfig.DB,
		Network:  "tcp",
	}
	return zoom.NewPoolWithOptions(options)
}

type Logger struct {
	Name string
}

// Module provided to fx
var Module = fx.Options(
	fx.Invoke(ConnectRedis),
	fx.Provide(ConnectRedis),
)
