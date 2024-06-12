package config

import (
	"log"
	"os"
	"time"
	"github.com/ilyakaznacheev/cleanenv"
)


type HTTPServerConfig struct {
	Host    string        `yaml:"host" env-default:"0.0.0.0"`
	Port    string        `yaml:"port" env-default:"8080"`
	Timeout time.Duration `yaml:"timeout" env-default:"5s"`
}

type Config struct {
	Env     string         `yaml:"env"`
	DbUrl   string         `yaml:"db_url" env-required:"true"`
	HTTP    HTTPServerConfig
}

func MustLoad() Config {
	configPath := ".\\config\\local.yaml" // TODO remove hardcode
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s doesn't exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("cannot read config file: ", err)
	}

	return cfg
}
