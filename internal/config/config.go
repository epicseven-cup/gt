package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Headers []string `json:"headers"`
	Tags    []string `json:"tags"`
}

func NewConfig() (*Config, error) {
	configData, err := os.ReadFile("~/.config/gt/config.json")
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = json.Unmarshal(configData, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
