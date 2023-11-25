package main

import (
	"bolt-watcher/bolt"
	"bolt-watcher/service"
	"bolt-watcher/storage"
	"bolt-watcher/watcher"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	ctx, cancelFn := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelFn()

	apiClient := bolt.New(os.Getenv("login"), os.Getenv("password"))

	store, err := storage.New(os.Getenv("db"))
	if err != nil {
		log.Fatalf("%+v", err)
	}

	service := service.New(apiClient)

	go watcher.
		New(watcher.Config{
			Service: service,
			Store:   store,
			Timeout: 30 * time.Second,
		}).
		Start(ctx)

	<-ctx.Done()
}
