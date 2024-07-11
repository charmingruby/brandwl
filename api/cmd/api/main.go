package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmingruby/brandwl/config"
	"github.com/charmingruby/brandwl/internal/domain/search/search_usecase"
	mongoRepo "github.com/charmingruby/brandwl/internal/infra/database/mongo"
	"github.com/charmingruby/brandwl/internal/infra/transport/rest"
	v1 "github.com/charmingruby/brandwl/internal/infra/transport/rest/endpoint/v1"
	mongoConn "github.com/charmingruby/brandwl/pkg/mongo"
	"github.com/charmingruby/brandwl/tests/fake"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
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

	router := gin.Default()

	initDependencies(db, router)

	server := rest.NewServer(router, cfg.ServerConfig.Port)

	go func() {
		if err := server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error(fmt.Sprintf("[REST SERVER] %s", err.Error()))
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	slog.Info("HTTP Server interruption received!")

	ctx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Server.Shutdown(ctx); err != nil {
		slog.Error(fmt.Sprintf("[GRACEFUL SHUTDOWN REST SERVER] %s", err.Error()))
		os.Exit(1)
	}

	slog.Info("Gracefully shutdown!")
}

func initDependencies(db *mongo.Database, router *gin.Engine) {
	searchRepo := mongoRepo.NewMongoSearchRepository(db)
	searchUseCase := search_usecase.NewSearchUseCase(searchRepo, fake.NewFakeGoogleAPI())
	v1.NewV1Handler(router, searchUseCase).Register()
}
