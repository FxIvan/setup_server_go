package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		HTTP  *HTTP
		Token *Token
	}

	Token struct {
		PublicKey  string
		PrivateKey string
		Duration   string
	}

	HTTP struct {
		Env            string
		URL            string
		Port           string
		AllowedOrigins string
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	http := &HTTP{
		Env:            os.Getenv("APP_ENV"),
		URL:            os.Getenv("HTTP_URL"),
		Port:           os.Getenv("HTTP_PORT"),
		AllowedOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
	}

	token := &Token{
		PrivateKey: os.Getenv("PRIVATE_KEY"),
		PublicKey:  os.Getenv("PUBLIC_KEY"),
		Duration:   os.Getenv("TOKEN_DURATION"),
	}

	return &Container{
		http,
		token,
	}, nil
}
