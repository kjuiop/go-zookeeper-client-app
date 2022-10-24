package util

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	ApiInfo struct {
		Port string `envconfig:"ZK_PORT" default:"3010"`
	}

	LogInfo struct {
		LogPath  string `envconfig:"ZK_LOG_PATH" default:"log/reporter.log"`
		LogLevel string `envconfig:"ZK_LOG_LEVEL" default:"debug"`
	}
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
