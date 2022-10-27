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

	ZookeeperInfo struct {
		EnsembleCount   string `envconfig:"ZK_ZOOKEPPER_ENSEMBLE_COUNT" default:"3"`
		Host            string `envconfig:"ZK_ZOOKEPPER_HOST" default:"172.21.0.1:2185"`
		RootNode        string `envconfig:"ZK_ZOOKEPPER_ROOT_NODE" default:"/zookeeper"`
		GroupName       string `envconfig:"ZK_ZOOKEEPER_GROUP_NAME" default:"tortee"`
		PollingInterval string `envconfig:"ZK_POLLING_INTERVAL" default:"10"`
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
