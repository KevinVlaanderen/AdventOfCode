package model

import "2023/src/framework/geometry"

type Segment struct {
	Pipe  *Pipe
	Point *geometry.Point
}

func (s Segment) DirectionOf(other Segment) Direction {
	switch {
	case other.Point.Y < s.Point.Y:
		return Top
	case other.Point.X > s.Point.X:
		return Right
	case other.Point.Y > s.Point.Y:
		return Bottom
	case other.Point.X < s.Point.X:
		return Left
	}
	panic("invalid direction")
}
