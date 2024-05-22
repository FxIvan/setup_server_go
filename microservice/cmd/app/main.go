package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	error_log *log.Logger
	info_log  *log.Logger
}

func main() {
	info_log := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	error_log := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	//go run . -srvPort 8000 -srvAddr localhost
	server_addr := flag.String("srvAddr", "", "HTTP server network address")
	server_port := flag.Int("srvPort", 4000, "HTTP server network port")
	flag.Parse()

	app := &application{
		error_log: error_log,
		info_log:  info_log,
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
