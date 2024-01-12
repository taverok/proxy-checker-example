package config

import (
	"os"

	"github.com/taverok/proxy-checker-example/pkg/db"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Name   string `yaml:"name"`
	Server struct {
		Port    int `yaml:"port"`
		Timeout int `yaml:"timeout"`
	} `yaml:"server"`
	DB db.Datasource `yaml:"db"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	cfg.Name = "checker"
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
