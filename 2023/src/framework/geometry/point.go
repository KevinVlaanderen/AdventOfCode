package geometry

import (
	"2023/src/framework"
)

type Point struct {
	X, Y int
}

func (p Point) Neighbors() (points []Point) {
	for x := range framework.RangeGen(p.X-1, 3, 1) {
		for y := range framework.RangeGen(p.Y-1, 3, 1) {
			if x != p.X || y != p.Y {
				points = append(points, Point{x, y})
			}
		}
	}
	return
}

func (p Point) Up() Point {
	return Point{p.X, p.Y - 1}
}

func (p Point) Right() Point {
	return Point{p.X + 1, p.Y}
}

func (p Point) Down() Point {
	return Point{p.X, p.Y + 1}
}

func (p Point) Left() Point {
	return Point{p.X - 1, p.Y}
}
