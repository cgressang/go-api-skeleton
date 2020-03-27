package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Log struct {
		Debug bool
	}
	Database struct {
		User     string
		Password string
		Net      string
		Host     string
		Port     int
		DbName   string
	}
	Server struct {
		Port string
	}
}

func New() (*Config, error) {
	config := &Config{}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join("$GOPATH", "src", "github.com", "cgressang", "go-api-skeleton", "deployment", os.Getenv("ENV")))

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
