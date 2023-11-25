package service

import (
	"bolt-watcher/bolt"
	"bolt-watcher/utils"
	"context"

	"github.com/samber/lo"
)

type Route []bolt.Point

type RouteCost struct {
	ETA        int64   // minutes.
	Price      int64   // hryvna.
	Multiplier float64 // X.
}

func (s *Service) GetRouteCosts(
	ctx context.Context,
	route Route,
) ([]*RouteCost, error) {
	res, err := s.API.GetRideOptions(ctx, &bolt.RideOptionsRequest{
		PickupStop:       route[0],
		DestinationStops: route[1:],
		PaymentMethod:    DefaultPayment,
	})
	if err != nil {
		return nil, err
	}

	resList := []*RouteCost{}

	taxis := res.Data.RideOptions[bolt.OrderSystemTaxi].Categories

	if taxis != nil {
		for _, category := range res.Data.CategoriesList {
			if category.OrderSystem != bolt.OrderSystemTaxi {
				continue
			}

			option, ok := taxis[category.CategoryID]
			if !ok {
				continue
			}

			if lo.IndexOf(AllowedCostGroups, option.Group) == -1 {
				continue
			}

			res := &RouteCost{
				ETA:        utils.TryParseInt(option.ETAInfo.PickupETAStr),
				Multiplier: utils.TryParseFloat(option.Price.SurgeStr),
			}

			res.Price, _ = lo.Coalesce(
				utils.TryParseInt(option.Price.SecondLineHTML),
				utils.TryParseInt(option.Price.FirstLineHTML),
				utils.TryParseInt(option.Price.ActualStr),
			)

			resList = append(resList, res)
		}
	}

	return resList, nil
}
