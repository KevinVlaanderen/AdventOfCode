package geo2d

import (
	"aoc/framework"
	"aoc/framework/math"
)

type Area struct {
	U, V Point
}

func NewArea(u, v Point) Area {
	return Area{
		Point{
			X: math.MinInt(u.X, v.X),
			Y: math.MinInt(u.Y, v.Y),
		},
		Point{
			X: math.MaxInt(u.X, v.X),
			Y: math.MaxInt(u.Y, v.Y),
		},
	}
}

func (a Area) Width() int {
	return 1 + math.AbsInt(a.U.X-a.V.X)
}

func (a Area) Height() int {
	return 1 + math.AbsInt(a.U.Y-a.V.Y)
}

func (a Area) Size() int {
	return a.Width() * a.Height()
}

func (a Area) Neighbors() (points []Point) {
	width, height := a.V.X-a.U.X+1, a.V.Y-a.U.Y+1
	for y := range framework.RangeGen(a.U.Y-1, height+2, 1) {
		points = append(points, Point{a.U.X - 1, y})
		points = append(points, Point{a.V.X + 1, y})
	}
	for x := range framework.RangeGen(a.U.X, width, 1) {
		points = append(points, Point{x, a.U.Y - 1})
		points = append(points, Point{x, a.V.Y + 1})
	}
	return
}

func (a Area) Contains(point Point) bool {
	return point.X >= a.U.X && point.X <= a.V.X && point.Y >= a.U.Y && point.Y <= a.V.Y
}
