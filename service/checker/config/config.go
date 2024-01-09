package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Name   string `yaml:"name"`
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	DB struct {
		Host string `yaml:"host"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Port int    `yaml:"port"`
		Name string `yaml:"name"`
	} `yaml:"db"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	path := "config/config.yml"

	env, ok := os.LookupEnv("DOMAIN_CHECK_CONFIG")
	if ok {
		path = env
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, cfg)
	if err != nil {
		return nil, err
	}

	env, ok = os.LookupEnv("DOMAIN_CHECK_DB_PASS")
	if ok {
		cfg.DB.Pass = env
	}

	return cfg, nil
}
