package geometry

import (
	"2023/src/framework/generators"
	"fmt"
)

type Grid[T comparable] struct {
	points                 map[Point]T
	xMin, xMax, yMin, yMax int
}

func NewGrid[T comparable]() Grid[T] {
	return Grid[T]{
		points: make(map[Point]T),
	}
}

func (g *Grid[T]) Get(key Point) (value T, found bool) {
	value, found = g.points[key]
	return
}

func (g *Grid[T]) Add(key Point, value T) {
	g.points[key] = value
	if len(g.points) == 0 || key.X < g.xMin {
		g.xMin = key.X
	}
	if len(g.points) == 0 || key.X > g.xMax {
		g.xMax = key.X
	}
	if len(g.points) == 0 || key.Y < g.yMin {
		g.yMin = key.Y
	}
	if len(g.points) == 0 || key.Y > g.yMax {
		g.yMax = key.Y
	}
}

func (g *Grid[T]) Neighbours(point Point) (neighbours []Point) {
	for x := point.X - 1; x <= point.X+1; x++ {
		for y := point.Y - 1; y <= point.Y+1; y++ {
			if x < g.xMin || x > g.xMax || y < g.yMin || y > g.yMax || (x == point.X && y == point.Y) {
				continue
			}
			neighbours = append(neighbours, Point{x, y})
		}
	}

	return
}

func (g *Grid[T]) NeighboursBy(point Point, filter func(neighbour Point) bool) (neighbours []Point) {
	for _, n := range g.Neighbours(point) {
		if filter(n) {
			neighbours = append(neighbours, n)
		}
	}

	return
}

func (g *Grid[T]) DrawPointGrid(mapping map[T]rune, fallback rune) {
	for y := range generators.RangeGen(g.yMin, g.yMax-g.yMin+1, 1) {
		for x := range generators.RangeGen(g.xMin, g.xMax-g.xMin+1, 1) {
			if value, positionExists := g.points[Point{X: x, Y: y}]; positionExists {
				if character, valueExists := mapping[value]; valueExists {
					fmt.Print(string(character))
				} else {
					fmt.Print(string(fallback))
				}
			} else {
				fmt.Print(string(fallback))
			}
		}
		fmt.Print("\n")
	}
}
