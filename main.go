package main

import (
	"bolt-watcher/api"
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

	apiClient := api.NewAPI(os.Getenv("login"), os.Getenv("password"))

	store, err := storage.New(os.Getenv("db"))
	if err != nil {
		log.Fatalf("%+v", err)
	}

	go watcher.
		New(watcher.Config{
			API:     apiClient,
			Store:   store,
			Timeout: 30 * time.Second,
		}).
		Start(ctx)

	<-ctx.Done()
}
