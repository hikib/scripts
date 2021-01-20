package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Pages map[string]string `yaml:pages`
}

func GetConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := yaml.NewDecoder(file)
	if err := data.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
