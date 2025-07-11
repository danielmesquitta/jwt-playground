package main

import (
	"context"
	"log"

	"github.com/danielmesquitta/jwt-playground/internal/app/server"
	"github.com/danielmesquitta/jwt-playground/internal/pkg/gracefulshutdown"
)

func main() {
	ctx := gracefulshutdown.WithShutdownSignal(
		context.Background(),
		logShutdown,
	)

	s := server.New()
	if err := s.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

func logShutdown() {
	log.Println("shutting down server")
}
