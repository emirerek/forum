package config

import (
	"encoding/json"
	"os"
)

type AdminConfig struct {
	Username string `json:"adminUsername"`
	Email    string `json:"adminEmail"`
	Password string `json:"adminPassword"`
}

type Config struct {
	Host   string      `json:"host"`
	Port   string      `json:"port"`
	DSN    string      `json:"dsn"`
	Secret string      `json:"secret"`
	Admin  AdminConfig `json:"admin"`
}

func Load(path string) (*Config, error) {
	config := &Config{}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
