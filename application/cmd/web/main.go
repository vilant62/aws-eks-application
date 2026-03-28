package main

import (
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {}

func main() {
	app := application{}
	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
	log.Println("Starting application on port", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

