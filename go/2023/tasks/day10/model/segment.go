package model

import (
	"aoc/framework/geometry/geo2d"
)

type Segment struct {
	Pipe  *Pipe
	Point *geo2d.Point
}

func (s Segment) DirectionOf(other Segment) geo2d.Orientation {
	switch {
	case other.Point.Y < s.Point.Y:
		return geo2d.North
	case other.Point.X > s.Point.X:
		return geo2d.East
	case other.Point.Y > s.Point.Y:
		return geo2d.South
	case other.Point.X < s.Point.X:
		return geo2d.West
	}
	panic("invalid direction")
}
