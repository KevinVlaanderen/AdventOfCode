package geometry

import (
	"2023/src/framework/generators"
	"fmt"
)

type Grid[T comparable] map[Point]T

func (g Grid[T]) Boundaries() (int, int, int, int) {
	initialized := false
	var xMin, xMax, yMin, yMax int
	for position := range g {
		if !initialized || position.X < xMin {
			xMin = position.X
		}
		if !initialized || position.X > xMax {
			xMax = position.X
		}
		if !initialized || position.Y < yMin {
			yMin = position.Y
		}
		if !initialized || position.Y > yMax {
			yMax = position.Y
		}
		initialized = true
	}
	return xMin, xMax, yMin, yMax
}

func (g Grid[T]) Neighbours(point Point) <-chan Point {
	c := make(chan Point)

	xMin, xMax, yMin, yMax := g.Boundaries()

	go func() {
		defer close(c)
		for x := point.X - 1; x <= point.X+1; x++ {
			for y := point.Y - 1; y <= point.Y+1; y++ {
				if x < xMin || x > xMax || y < yMin || y > yMax || (x == point.X && y == point.Y) {
					continue
				}
				c <- Point{x, y}
			}
		}
	}()

	return c
}

func (g Grid[T]) NeighboursBy(point Point, filter func(neighbour Point) bool) <-chan Point {
	c := make(chan Point)

	go func() {
		defer close(c)
		for n := range g.Neighbours(point) {
			if filter(n) {
				c <- n
			}
		}
	}()

	return c
}

func (g Grid[T]) DrawPointGrid(mapping map[T]rune, fallback rune) {
	xMin, xMax, yMin, yMax := g.Boundaries()

	for y := range generators.RangeGen(yMin, yMax-yMin+1, 1) {
		for x := range generators.RangeGen(xMin, xMax-xMin+1, 1) {
			if value, positionExists := g[Point{X: x, Y: y}]; positionExists {
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
