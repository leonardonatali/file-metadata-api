package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Port       int    `envconfig:"PORT" default:"80"`
	DBPort     int    `envconfig:"DB_PORT" default:"27018"`
	DBHost     string `envconfig:"DB_HOST" default:"database"`
	DBUser     string `envconfig:"DB_USER" default:"root"`
	DBName     string `envconfig:"DB_NAME" default:"files"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"root"`
}

func (c *Config) Load() error {
	return envconfig.Process("", c)
}

func (c *Config) GetDatabaseDSN(direct bool) string {

	if direct {
		return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
	}

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort,
	)
}

func (c *Config) GetDBConn() *gorm.DB {
	db, err := gorm.Open(postgres.Open(c.GetDatabaseDSN(false)), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: false,
				Colorful:                  true,
			},
		),
	})

	if err != nil {
		panic(err)
	}

	return db
}
