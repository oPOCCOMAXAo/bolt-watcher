package main

import (
	"bolt-watcher/api"
	"bolt-watcher/logger"
	"bolt-watcher/spreadsheet"
	"bolt-watcher/watcher"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	apiClient := api.NewAPI(os.Getenv("login"), os.Getenv("password"))

	c := make(chan os.Signal, 10)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	gsheet, err := spreadsheet.New("./key.json")
	if err != nil {
		panic(err)
	}

	workPos := api.Point{
		Latitude:  50.45035189213392,
		Longitude: 30.44201365424139,
	}
	homePos := api.Point{
		Latitude:  50.48144271866205,
		Longitude: 30.395991111470966,
	}

	go watcher.
		New(watcher.Config{
			API:  apiClient,
			From: workPos,
			To:   homePos,
			Logger: logger.New(logger.Config{
				Service:       gsheet,
				SpreadsheetId: "13cOT9QjGLHzx_zETWE7SGnfNBS5dTAaa5hbBtX0cEZE",
				SheetId:       "work-home",
			}),
			Timeout: time.Minute * 2,
		}).
		Start()

	go watcher.
		New(watcher.Config{
			API:  apiClient,
			From: homePos,
			To:   workPos,
			Logger: logger.New(logger.Config{
				Service:       gsheet,
				SpreadsheetId: "13cOT9QjGLHzx_zETWE7SGnfNBS5dTAaa5hbBtX0cEZE",
				SheetId:       "home-work",
			}),
			Timeout: time.Minute * 2,
		}).
		Start()

	<-c
}
