package main

import (
	"context"
	"go-enum-example/app"
	"log"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	shutdown := app.Run(ctx)

	<-ctx.Done() // wait until signal received
	log.Println("shutting down server")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	shutdown(shutdownCtx)
}
