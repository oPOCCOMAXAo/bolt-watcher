package watcher

import (
	"bolt-watcher/bolt"
	"bolt-watcher/service"
	"bolt-watcher/storage"
	"bolt-watcher/utils"
	"context"
	"log"
	"math"
	"time"

	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type Watcher struct {
	cfg Config
}

type Config struct {
	Service *service.Service
	Store   *storage.Storage
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

func (w *Watcher) Start(ctx context.Context) {
	ticker := time.NewTicker(w.cfg.Timeout)

	done := ctx.Done()

	w.checkErr(w.tick(ctx))

	for {
		select {
		case <-ticker.C:
			w.checkErr(w.tick(ctx))
		case <-done:
			return
		}
	}
}

func (w *Watcher) checkErr(err error) {
	if err != nil {
		log.Printf("%+v\n", err)
	}
}

func (w *Watcher) tick(ctx context.Context) error {
	routes, err := w.cfg.Store.GetActiveRoutes(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	for _, route := range routes {
		w.checkErr(w.processRoute(ctx, route))
	}

	return nil
}

func (w *Watcher) processRoute(ctx context.Context, route *storage.RouteExt) error {
	res, err := w.cfg.Service.GetRouteCosts(ctx,
		lo.Map(route.Route, func(r storage.PointExt, _ int) bolt.Point {
			return bolt.Point{Lat: r.Lat, Lon: r.Lon}
		}),
	)
	if err != nil {
		return errors.WithStack(err)
	}

	if len(res) > 0 {
		return w.cfg.Store.Log(ctx, storage.Log{
			Time:       utils.Floor(time.Now().Unix(), 60),
			RouteID:    route.ID,
			ETA:        res[0].ETA,
			Price:      res[0].Price,
			Multiplier: math.Max(res[0].Multiplier, 1),
		})
	}

	return errors.WithStack(ErrNoResults)
}
