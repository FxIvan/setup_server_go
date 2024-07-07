package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		HTTP            *HTTP
		Token           *Token
		JWT             *JWT
		TerminalLog     *TerminalLog
		URLMicroservice *URLMicroservice
	}

	Token struct {
		PublicKey  string
		PrivateKey string
		Duration   string
	}

	JWT struct {
		JWT_SCRET string
		DURATION  string
	}

	HTTP struct {
		Env            string
		URL            string
		Port           string
		AllowedOrigins string
	}

	TerminalLog struct {
		ErrorLog *log.Logger
		InfoLog  *log.Logger
	}

	URLMicroservice struct {
		HostCreatePaymentNodeJS string
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println("Error loading .env file", err)
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

	terminalLog := &TerminalLog{
		ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		InfoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	}

	jwt := &JWT{
		JWT_SCRET: os.Getenv("JWT_SECRET"),
		DURATION:  os.Getenv("JWT_DURATION"),
	}

	urlMicroservice := &URLMicroservice{
		HostCreatePaymentNodeJS: os.Getenv("URL_CREATE_PAYMENT_NODEJS"),
	}

	return &Container{
		http,
		token,
		jwt,
		terminalLog,
		urlMicroservice,
	}, nil
}
