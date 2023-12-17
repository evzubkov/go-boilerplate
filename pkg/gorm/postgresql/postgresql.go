package postgresql

import (
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	Logger   logger.Interface
}

// NewPostgresClient - establishes a connection to a PostgreSQL database using the provided configuration.
func NewPostgresClient(config DbConfig) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName)

	if config.Logger == nil {
		config.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		)
	}

	// Connect to the PostgreSQL database
	if db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: config.Logger}); err != nil {
		return
	}

	return
}
