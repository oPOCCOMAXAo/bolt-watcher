package watcher

import (
	"bolt-watcher/api"
	"time"
)

type Watcher struct {
	cfg Config
}

type Config struct {
	API     *api.API
	From    api.Point
	To      api.Point
	Logger  Logger
	Timeout time.Duration
}

func New(config Config) *Watcher {
	if config.Timeout == 0 {
		config.Timeout = time.Minute
	}
	return &Watcher{
		cfg: config,
	}
}

func (w *Watcher) Start() {
	for {
		w.tick()
		time.Sleep(w.cfg.Timeout)
	}
}

func (w *Watcher) tick() {
	res := w.cfg.API.FindCategoriesOverview(api.FindCategoriesOverviewRequest{
		From: w.cfg.From,
		To:   w.cfg.To,
	})

	var c *api.SearchCategory
	for i, cat := range res.Data.SearchCategories {
		if cat.Name == "Bolt" {
			c = &res.Data.SearchCategories[i]
			break
		}
	}

	if c == nil {
		c = &res.Data.SearchCategories[0]
	}

	if c == nil {
		return
	}

	w.cfg.Logger.Log(time.Now(), parsePrice(c.UpfrontPriceStr), c.SurgeMultiplier)
}
