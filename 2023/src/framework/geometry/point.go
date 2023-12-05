package geometry

import "2023/src/framework/generators"

type Point struct {
	X, Y int
}

func (p Point) Neighbors() (points []Point) {
	for _, x := range generators.Range(p.X-1, 3, 1) {
		for _, y := range generators.Range(p.Y-1, 3, 1) {
			if x != p.X && y != p.Y {
				points = append(points, Point{x, y})
			}
		}
	}
	return
}
