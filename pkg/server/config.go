package server

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Port    int
	Postgre PostgreConfig
}

type PostgreConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
	Timeout  int
}

func ReadConfig(path string) (Config, error) {
	var cfg Config
	_, err := toml.DecodeFile(path, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
