package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/charmingruby/brandwl/config"
	"github.com/charmingruby/brandwl/pkg/mongo"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Info("[CONFIGURATION] .env not found.")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("[CONFIGURATION] %s", err.Error()))
		os.Exit(1)
	}

	_, err = mongo.NewMongoConnection(cfg.MongoConfig.URL, cfg.MongoConfig.Database)
	if err != nil {
		slog.Error(fmt.Sprintf("[MONGO CONNECTION] %s", err.Error()))
		os.Exit(1)
	}
}
