package grid

import (
	"2023/src/framework"
	"2023/src/framework/geometry"
	"fmt"
	"github.com/samber/lo"
)

type SparseGrid[T comparable] struct {
	points                 map[geometry.Point]T
	xMin, xMax, yMin, yMax int
}

func NewSparseGrid[T comparable]() SparseGrid[T] {
	return SparseGrid[T]{
		points: make(map[geometry.Point]T),
	}
}

func (g *SparseGrid[T]) Keys() []geometry.Point {
	return lo.Keys(g.points)
}

func (g *SparseGrid[T]) Has(key geometry.Point) (found bool) {
	_, found = g.Get(key)
	return
}

func (g *SparseGrid[T]) Get(key geometry.Point) (value T, found bool) {
	value, found = g.points[key]
	return
}

func (g *SparseGrid[T]) Add(key geometry.Point, value T) {
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

func (g *SparseGrid[T]) Delete(key geometry.Point) {
	delete(g.points, key)
}

func (g *SparseGrid[T]) Boundaries() (int, int, int, int) {
	return g.xMin, g.xMax, g.yMin, g.yMax
}

func (g *SparseGrid[T]) DrawPointGrid(mapping map[T]rune, fallback rune) {
	g.DrawPointGridBy(func(value T, found bool, x int, y int) rune {
		if !found {
			return fallback
		} else if character, valueExists := mapping[value]; valueExists {
			return character
		} else {
			return fallback
		}
	})
}

func (g *SparseGrid[T]) DrawPointGridBy(mapping func(value T, found bool, x int, y int) rune) {
	for y := range framework.RangeGen(g.yMin, g.yMax-g.yMin+1, 1) {
		for x := range framework.RangeGen(g.xMin, g.xMax-g.xMin+1, 1) {
			value, found := g.points[geometry.Point{X: x, Y: y}]

			if mapping != nil {
				fmt.Print(string(mapping(value, found, x, y)))
				continue
			}
		}
		fmt.Print("\n")
	}
}

func (g *SparseGrid[T]) InBounds(point geometry.Point) bool {
	return point.X >= g.xMin && point.X <= g.xMax && point.Y >= g.yMin && point.Y <= g.yMax
}
