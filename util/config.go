package util

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	ApiPort string `envconfig:"API_PORT" default:"3010"`
}

func ConfInitialize() (*Config, error) {
	c := new(Config)

	err := envconfig.Process("api", c)
	if err != nil {
		log.Println("[ConfInitialize] failed read config :", err)
		return nil, err
	}

	return c, nil
}
