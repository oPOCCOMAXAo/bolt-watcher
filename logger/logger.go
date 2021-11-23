package logger

import (
	"bolt-watcher/spreadsheet"
	"fmt"
	"time"
)

type Logger struct {
	cfg Config
}

type Config struct {
	Service       *spreadsheet.Service
	SpreadsheetId string
	SheetId       string
}

func New(cfg Config) *Logger {
	return &Logger{
		cfg: Config{
			Service: cfg.Service.WithSpreadsheet(cfg.SpreadsheetId).WithSheet(cfg.SheetId),
		},
	}
}

func (l *Logger) Log(stamp time.Time, price float64, mult float64) {
	err := l.cfg.Service.InsertRow(
		stamp.Format("02.01.2006 15:04:05"),
		price,
		mult,
	)
	if err != nil {
		fmt.Println(err)
	}
}
