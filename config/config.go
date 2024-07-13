package config

import (
	"log/slog"

	env "github.com/caarlos0/env/v6"
)

type environment struct {
	MongoURL       string `env:"MONGO_URL,required"`
	MongoDatabase  string `env:"MONGO_DATABASE,required"`
	ServerPort     string `env:"PORT,required"`
	MailerHost     string `env:"MAILER_HOST,required"`
	MailerPort     int    `env:"MAILER_PORT,required"`
	MailerFrom     string `env:"MAILER_FROM,required"`
	MailerUser     string `env:"MAILER_USER,required"`
	MailerPassword string `env:"MAILER_PASSWORD,required"`
	GoogleAPIKey   string `env:"GOOGLE_API_KEY,required"`
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
		MailerConfig: &mailerConfig{
			Host:     environment.MailerHost,
			Port:     environment.MailerPort,
			From:     environment.MailerFrom,
			User:     environment.MailerUser,
			Password: environment.MailerPassword,
		},
		GoogleConfig: &googleConfig{
			APIKey: environment.GoogleAPIKey,
		},
	}

	return &cfg, nil
}

type Config struct {
	MongoConfig  *mongoConfig
	ServerConfig *serverConfig
	MailerConfig *mailerConfig
	GoogleConfig *googleConfig
}

type mongoConfig struct {
	Database string
	URL      string
}

type mailerConfig struct {
	Host     string
	Port     int
	From     string
	User     string
	Password string
}

type googleConfig struct {
	APIKey string
}

type serverConfig struct {
	Port string
}
