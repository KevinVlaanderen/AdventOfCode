package geometry

import "2023/src/framework/generators"

type Point struct {
	X, Y int
}

func (p Point) Neighbors() (points []Point) {
	for x := range generators.RangeGen(p.X-1, 3, 1) {
		for y := range generators.RangeGen(p.Y-1, 3, 1) {
			if x != p.X || y != p.Y {
				points = append(points, Point{x, y})
			}
		}
	}
	return
}
