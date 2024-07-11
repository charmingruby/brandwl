package config

import (
	"log/slog"

	env "github.com/caarlos0/env/v6"
)

type environment struct {
	MongoURL      string `env:"MONGO_URL,required"`
	MongoDatabase string `env:"MONGO_DATABASE,required"`
	ServerPort    string `env:"PORT,required"`
}

func NewConfig() (*Config, error) {
	slog.Info("Loading environment...")

	environment := environment{}
	if err := env.Parse(&environment); err != nil {
		return nil, err
	}

	slog.Info("Environment loaded successfully!")

	cfg := Config{
		MongoConfig: &mongoConfig{
			Database: environment.MongoDatabase,
			URL:      environment.MongoURL,
		},
		ServerConfig: &serverConfig{
			Port: environment.ServerPort,
		},
	}

	return &cfg, nil
}

type Config struct {
	MongoConfig  *mongoConfig
	ServerConfig *serverConfig
}

type mongoConfig struct {
	Database string
	URL      string
}

type serverConfig struct {
	Port string
}