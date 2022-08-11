package config

import (
	"smartHome/pkg/logging"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env:"SHDebug"`
	App     struct {
		Host string `yaml:"host" env:"SHHost" env-default:"localhost"`
		Port int    `yaml:"port" env:"SHPort" env-default:"50194"`
	}
	MongoDB struct {
		HostMDB    string `yaml:"hostmdb" env:"SHHostMDB" env-default:"localhost"`
		PortMDB    string `yaml:"portmdb" env:"SHPortMDB" env-default:"50195"`
		Database   string `yaml:"database" env:"SHdatabaseMDB" env-default:"userservice"`
		AuthDB     string `yaml:"host" env:"SHAuthDBMDB"`
		Collection string `yaml:"collection" env:"SHCollectionMDB" env-default:"users"`
		Username   string `yaml:"username" env:"SHUsernameMDB"`
		Password   string `yaml:"password" env:"SHPasswordMDB"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read Config")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yaml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
