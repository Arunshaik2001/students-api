package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type HttpServerConfig struct {
	Addr string `yaml:"address"`
}

type Config struct {
	Env         string           `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string           `yaml:"storage_path" env-required:"true"`
	HttpServer  HttpServerConfig `yaml:"http_server"`
}

func LoadConfig() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to config file")

		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("No Config File Provided")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config File Not Found %s", err)
	}

	var config Config

	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		log.Fatalf("Error %s", err.Error())
	}

	return &config
}
