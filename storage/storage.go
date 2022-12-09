package storage

import (
	"context"
	"log"
	"sort"
	"time"

	"github.com/opoccomaxao-go/generic-collection/slice"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type Storage struct {
	db *gorm.DB
}

func New(connection string) (*Storage, error) {
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{
		Logger: logger.New(log.Default(), logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Error,
		}),
	})
	if err != nil {
		return nil, err
	}

	res := Storage{
		db: db,
	}

	return &res, res.init()
}

func (s *Storage) GetActiveRoutes(
	ctx context.Context,
) ([]*RouteExt, error) {
	var res []*RoutePointCombo

	err := s.db.WithContext(ctx).Table("route_point AS rp").
		Select(
			"rp.*",
			"p.lat",
			"p.lon",
		).
		Joins("JOIN route AS r ON rp.route_id = r.id AND r.active").
		Joins("JOIN point AS p ON p.id = rp.point_id").
		Find(&res).
		Error
	if err != nil {
		return nil, err
	}

	splitRoutes := map[int][]*RoutePointCombo{}
	for _, rp := range res {
		splitRoutes[rp.RouteID] = append(splitRoutes[rp.RouteID], rp)
	}

	resRoutes := []*RouteExt{}
	for k, v := range splitRoutes {
		pts := v

		sort.Slice(pts, func(i, j int) bool {
			return pts[i].Position < pts[j].Position
		})

		resRoutes = append(resRoutes, &RouteExt{
			ID: k,
			Route: slice.Map(pts, func(p *RoutePointCombo) PointExt {
				return PointExt{
					Lat: p.Lat,
					Lon: p.Lon,
				}
			}),
		})
	}

	return resRoutes, nil
}

func (s *Storage) Log(ctx context.Context, log Log) error {
	return s.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			DoUpdates: clause.Assignments(map[string]interface{}{
				"eta":        gorm.Expr("GREATEST(VALUES(eta),eta)"),
				"price":      gorm.Expr("GREATEST(VALUES(price),price)"),
				"multiplier": gorm.Expr("GREATEST(VALUES(multiplier),multiplier)"),
			}),
		}).
		Create(&log).
		Error
}
