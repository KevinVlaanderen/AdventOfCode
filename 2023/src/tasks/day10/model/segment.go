package model

import "2023/src/framework/geometry"

type Segment struct {
	Pipe  *Pipe
	Point *geometry.Point
}

func (s Segment) DirectionOf(other Segment) geometry.Orientation {
	switch {
	case other.Point.Y < s.Point.Y:
		return geometry.North
	case other.Point.X > s.Point.X:
		return geometry.East
	case other.Point.Y > s.Point.Y:
		return geometry.South
	case other.Point.X < s.Point.X:
		return geometry.West
	}
	panic("invalid direction")
}
