package watcher

import "time"

type Logger interface {
	Log(stamp time.Time, price float64, mult float64)
}
