package model

import (
	"2022/src/framework/geometry"
	"2022/src/framework/math"
)

type Sensor struct {
	geometry.Point
	ClosestBeacon geometry.Point
	Distance int
}

func NewSensor(point geometry.Point, closestBeacon geometry.Point) Sensor {
	diffX := math.AbsInt(point.X - closestBeacon.X)
	diffY := math.AbsInt(point.Y - closestBeacon.Y)
	distance := diffX + diffY
	
	return Sensor{
		Point:         point,
		ClosestBeacon: closestBeacon,
		Distance:      distance,
	}
}

func (s Sensor) DistanceToBeacon() int {
	return math.AbsInt(s.X-s.ClosestBeacon.X)+math.AbsInt(s.Y-s.ClosestBeacon.Y)
}