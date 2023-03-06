package dbfx

import (
	"fmt"
	"log"

	"github.com/veteran-dev/veteran/pkg/system/configfx"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB(c *configfx.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", c.PostgresHost, c.PostgresUser, c.PostgresPassword, c.PostgresDbName, c.PostgresPort, c.PostgresSslMode, c.PostgresTimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	return db
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(ConnectDB),
)
