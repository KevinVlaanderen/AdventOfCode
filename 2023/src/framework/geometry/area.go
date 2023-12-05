package geometry

import "2023/src/framework/generators"

type Area struct {
	From, To Point
}

func (a Area) Neighbors() (points []Point) {
	width, height := a.To.X-a.From.X+1, a.To.Y-a.From.Y+1
	for _, y := range generators.Range(a.From.Y-1, height+2, 1) {
		points = append(points, Point{a.From.X - 1, y})
		points = append(points, Point{a.To.X + 1, y})
	}
	for _, x := range generators.Range(a.From.X, width, 1) {
		points = append(points, Point{x, a.From.Y - 1})
		points = append(points, Point{x, a.To.Y + 1})
	}
	return
}
