package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`

	Postgres struct {
		Username     string `yaml:"username"`
		DatabaseName string `yaml:"database_name"`
		Password     string `yaml:"password"`
		SslMode      string `yaml:"ssl_mode"`
	} `yaml:"postgres"`

	NatsServer struct {
		Url string `yaml:"url"`
	} `yaml:"nats_server"`
}

func ParseConfig(cfg *Config) error {
	configFile, err := os.Open("config/config.yml")
	if err != nil {
		return err
	}
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			logrus.Fatal(err)
		}
	}(configFile)
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}
	return nil
}
