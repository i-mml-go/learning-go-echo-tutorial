package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port string `yaml: "server"`
	} `yaml: "server"`
}

var AppConfig Config

func GetConfig() error {
	file, err := os.Open("config.yml")

	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	decoder.Decode(&AppConfig)
	if err != nil {
		return err
	}

	return nil
}
