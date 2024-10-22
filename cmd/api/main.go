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
	"github.com/charmingruby/brandwl/internal/domain/search/search_adapter"
	"github.com/charmingruby/brandwl/internal/domain/search/search_usecase"
	mongoRepo "github.com/charmingruby/brandwl/internal/infra/database/mongo"
	"github.com/charmingruby/brandwl/internal/infra/transport/rest"
	v1 "github.com/charmingruby/brandwl/internal/infra/transport/rest/endpoint/v1"
	"github.com/charmingruby/brandwl/pkg/mailer"
	mongoConn "github.com/charmingruby/brandwl/pkg/mongo"
	"github.com/charmingruby/brandwl/pkg/scrap"
	"github.com/gin-contrib/cors"
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

	mailer := mailer.NewMailTrapMailer(
		cfg.MailerConfig.Host,
		cfg.MailerConfig.Port,
		cfg.MailerConfig.From,
		cfg.MailerConfig.User,
		cfg.MailerConfig.Password,
	)

	scrap := scrap.NewGoogleSearch(cfg.GoogleConfig.APIKey)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Origin", "Accept", "Content-Type", "Authorization", "User-Agent"},
		ExposeHeaders:   []string{"Content-Length"},
	}))

	initDependencies(db, router, mailer, scrap)

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

func initDependencies(db *mongo.Database, router *gin.Engine, mailer mailer.Mailer, scrap search_adapter.GoogleAPIAdapter) {
	searchRepo := mongoRepo.NewMongoSearchRepository(db)
	searchUseCase := search_usecase.NewSearchUseCase(searchRepo, scrap)
	v1.NewV1Handler(router, searchUseCase, mailer).Register()
}
