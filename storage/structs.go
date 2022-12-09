package storage

type Point struct {
	ID   int     `gorm:"column:id;primaryKey;autoIncrement"`
	Name string  `gorm:"column:name;type:VARCHAR(256);not null"`
	Lat  float64 `gorm:"column:lat;not null"`
	Lon  float64 `gorm:"column:lon;not null"`
}

func (Point) TableName() string {
	return "point"
}

type Route struct {
	ID     int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name   string `gorm:"column:name;type:VARCHAR(256);not null"`
	Active bool   `gorm:"column:active;default:0;not null"`
}

func (Route) TableName() string {
	return "route"
}

type RoutePoint struct {
	RouteID  int `gorm:"column:route_id;primaryKey"`
	Position int `gorm:"column:position;primaryKey"`
	PointID  int `gorm:"column:point_id;not null"`
}

func (RoutePoint) TableName() string {
	return "route_point"
}

type Log struct {
	Time       int64   `gorm:"column:time;primaryKey"`
	RouteID    int     `gorm:"column:route_id;primaryKey"`
	ETA        int64   `gorm:"column:eta;not null"`
	Price      int64   `gorm:"column:price;not null"`
	Multiplier float64 `gorm:"column:multiplier;not null"`
}

func (Log) TableName() string {
	return "log"
}

type PointExt struct {
	Lat, Lon float64
}

type RouteExt struct {
	ID    int
	Route []PointExt
}

type RoutePointCombo struct {
	RoutePoint
	Point
}
