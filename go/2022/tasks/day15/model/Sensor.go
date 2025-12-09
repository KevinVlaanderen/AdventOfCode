package model

import (
	"aoc/framework/geometry/geo2d"
	"aoc/framework/math"
)

type Sensor struct {
	geo2d.Point
	ClosestBeacon geo2d.Point
	Distance      int
}

func NewSensor(point geo2d.Point, closestBeacon geo2d.Point) Sensor {
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
	return math.AbsInt(s.X-s.ClosestBeacon.X) + math.AbsInt(s.Y-s.ClosestBeacon.Y)
}
