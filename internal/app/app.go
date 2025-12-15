package app

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	}
	Database struct {
		DSN string `yaml:"dsn"`
	}
	Security struct {
		AdminUsername string `yaml:"admin_username"`
		AdminPassword string `yaml:"admin_password"`
	}
	Storage struct {
		DefaultPath string `yaml:"default_path"`
	}
}

func LoadConfig() *Config {
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}
