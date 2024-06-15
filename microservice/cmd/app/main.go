package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fxivan/set_up_server/microservice/configuration"
	"github.com/fxivan/set_up_server/microservice/internal/adapter/auth/paseto"
	"github.com/fxivan/set_up_server/microservice/internal/adapter/config"
	"github.com/fxivan/set_up_server/microservice/internal/adapter/handler/http"
	dbRepository "github.com/fxivan/set_up_server/microservice/internal/adapter/storage"
	service "github.com/fxivan/set_up_server/microservice/internal/core/service"
)

type ConnectDB struct {
	Engine   *string
	Host     *string
	Port     *int
	User     *string
	Password *string
	DBName   *string
}

type application struct {
	error_log *log.Logger
	info_log  *log.Logger
	connDB    *ConnectDB
}

func main() {
	info_log := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	error_log := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	//go run . -srvPort 8000 -srvAddr localhost
	//server_addr := flag.String("srvAddr", "", "HTTP server network address")
	//server_port := flag.Int("srvPort", 4000, "HTTP server network port")
	input_engine := flag.String("enginedb", "", "Engine DB")
	input_host := flag.String("hostdb", "localhost", "Host DB") //mongodb-container
	input_port := flag.Int("portdb", 3360, "Port DB")
	input_user := flag.String("userdb", "", "User DB")
	input_password := flag.String("passwordb", "", "Password DB")
	input_dbname := flag.String("dbname", "", "Name DB")
	flag.Parse()

	config, err := config.New()
	if err != nil {
		error_log.Println(err)
		os.Exit(1)
	}

	configDB := configuration.Configuration{
		Engine:   *input_engine,
		Host:     *input_host,
		Port:     *input_port,
		User:     *input_user,
		Password: *input_password,
		DBName:   *input_dbname,
	}

	repo, err := dbRepository.New(&configDB, config.TerminalLog)
	if err != nil {
		panic(err)
	}

	token, err := paseto.New(config.Token, config.TerminalLog)
	if err != nil {
		error_log.Println(err)
		os.Exit(1)
	}

	//User
	userService := service.NewUserService(repo, config.TerminalLog)
	userHandler := http.NewUserHandler(userService)

	//Auth
	authService := service.NewAuthService(repo, token, config.TerminalLog)
	authHandler := http.NewAuthHandler(authService)

	//GiftCard
	giftCardService := service.NewGiftCardService(repo, config.TerminalLog)
	giftCardHandler := http.NewGiftCardHandler(giftCardService)

	router, err := http.NewRouter(
		config.HTTP,
		token,
		*userHandler,
		*authHandler,
		*giftCardHandler,
	)

	if err != nil {
		error_log.Println(err)
		os.Exit(1)
	}

	//listenAddr := fmt.Sprintf(":%s", config.HTTP.Port)
	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	info_log.Printf("Starting server on %s", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		error_log.Println(err)
		os.Exit(1)
	}
}
