package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/charmingruby/brandwl/config"
	"github.com/charmingruby/brandwl/internal/domain/search/search_usecase"
	"github.com/charmingruby/brandwl/internal/infra/database/mongo"
	mongoConn "github.com/charmingruby/brandwl/pkg/mongo"
	"github.com/charmingruby/brandwl/tests/fake"
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

	db, err := mongoConn.NewMongoConnection(cfg.MongoConfig.URL, cfg.MongoConfig.Database)
	if err != nil {
		slog.Error(fmt.Sprintf("[MONGO CONNECTION] %s", err.Error()))
		os.Exit(1)
	}

	searchRepo := mongo.NewMongoSearchRepository(db)
	search_usecase.NewSearchUseCase(searchRepo, fake.NewFakeGoogleAPI())
}
