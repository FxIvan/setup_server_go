package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
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
	server_addr := flag.String("srvAddr", "", "HTTP server network address")
	server_port := flag.Int("srvPort", 4000, "HTTP server network port")
	input_engine := flag.String("enginedb", "", "Engine DB")
	input_host := flag.String("hostdb", "localhost", "Host DB")
	input_port := flag.Int("portdb", 3360, "Port DB")
	input_user := flag.String("userdb", "", "User DB")
	input_password := flag.String("passwordb", "", "Password DB")
	input_dbname := flag.String("dbname", "", "Name DB")
	flag.Parse()

	conn_obj := &ConnectDB{
		Engine:   input_engine,
		Host:     input_host,
		Port:     input_port,
		User:     input_user,
		Password: input_password,
		DBName:   input_dbname,
	}

	app := &application{
		error_log: error_log,
		info_log:  info_log,
		connDB:    conn_obj,
	}

	serverURI := fmt.Sprintf("%s:%d", *server_addr, *server_port)

	srv := &http.Server{
		Addr:         serverURI,
		ErrorLog:     error_log,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	info_log.Printf("Starting server on %s", serverURI)
	err := srv.ListenAndServe()
	error_log.Fatal(err)
}
