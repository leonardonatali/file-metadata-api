package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port       int    `envconfig:"port" default:"80"`
	DBPort     int    `envconfig:"db_port" default:"27018"`
	DBhost     string `envconfig:"db_bhost" default:"database"`
	DBUser     string `envconfig:"db_user" default:"root"`
	DBPassword string `envconfig:"db_password" default:"root"`
}

func (c *Config) Load() error {
	return envconfig.Process("", c)
}
