package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port       int    `envconfig:"port" default:"80"`
	DBPort     int    `envconfig:"db_port" default:"27018"`
	DBHost     string `envconfig:"db_bhost" default:"database"`
	DBUser     string `envconfig:"db_user" default:"root"`
	DBName     string `envconfig:"db_name" default:"files"`
	DBPassword string `envconfig:"db_password" default:"root"`
}

func (c *Config) Load() error {
	return envconfig.Process("", c)
}

func (c *Config) GetDatabaseDSN(direct bool) string {

	if direct {
		return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
	}

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Ameriaca/Sao_Paulo",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort,
	)
}
