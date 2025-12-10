package geo2d

import (
	"aoc/framework/math"
	"errors"
)

type Line struct {
	A, B Point
}

func (line Line) LineType() LineType {
	switch {
	case line.A.X == line.B.X && line.A.Y == line.B.Y:
		return Undefined
	case line.A.X == line.B.X:
		return Vertical
	case line.A.Y == line.B.Y:
		return Horizontal
	default:
		return Diagonal
	}
}

func (line Line) Direction() (Direction, error) {
	lineType := line.LineType()
	if lineType == Horizontal {
		if line.A.X < line.B.X {
			return Right, nil
		}
		return Left, nil
	} else if lineType == Vertical {
		if line.A.Y < line.B.Y {
			return Down, nil
		}
		return Up, nil
	}
	return 0, errors.New("direction can only be determined for horizontal and vertical lines")
}

func (line Line) Points() ([]Point, error) {
	lineType := line.LineType()
	if lineType != Horizontal && lineType != Vertical {
		return nil, errors.New("points can only be determined for horizontal and vertical lines")
	}

	points := make([]Point, 0)

	for x := math.MinInt(line.A.X, line.B.X); x <= math.MaxInt(line.A.X, line.B.X); x++ {
		for y := math.MinInt(line.A.Y, line.B.Y); y <= math.MaxInt(line.A.Y, line.B.Y); y++ {
			points = append(points, Point{X: x, Y: y})
		}
	}

	return points, nil
}

type LineType uint8

const (
	Undefined LineType = iota
	Horizontal
	Vertical
	Diagonal
)
