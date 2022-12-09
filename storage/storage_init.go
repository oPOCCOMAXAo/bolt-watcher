package storage

func (s *Storage) init() error {
	return s.db.Migrator().
		AutoMigrate(
			&Point{},
			&Route{},
			&RoutePoint{},
			&Log{},

			// refs.
			&struct {
				RoutePoint
				Route *Route `gorm:"foreignKey:route_id"`
				Point *Point `gorm:"foreignKey:point_id"`
			}{},

			&struct {
				Log
				Route *Route `gorm:"foreignKey:route_id"`
			}{},
		)
}
