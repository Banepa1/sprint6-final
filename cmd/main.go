package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.Default()

	server := server.NewServer(logger)

	if err := server.HttpServer.ListenAndServe(); err != nil {
		server.Logger.Fatal(err)
		return
	}
}
