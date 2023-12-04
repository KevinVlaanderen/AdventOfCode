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

func (g Grid[T]) DrawPointGrid(mapping map[T]rune, fallback rune) {
	xMin, xMax, yMin, yMax := g.Boundaries()

	for _, y := range generators.Range(yMin, yMax-yMin+1, 1) {
		for _, x := range generators.Range(xMin, xMax-xMin+1, 1) {
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
