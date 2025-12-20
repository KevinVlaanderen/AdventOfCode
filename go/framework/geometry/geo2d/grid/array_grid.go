package grid

import (
	"aoc/framework/geometry/geo2d"
	"fmt"

	"github.com/samber/lo"
)

type item[T comparable] struct {
	value T
	isSet bool
}

type ArrayGrid[T comparable] struct {
	data          []item[T]
	width, height int
}

func NewArrayGrid[T comparable](width, height int) *ArrayGrid[T] {
	data := make([]item[T], width*height)
	return &ArrayGrid[T]{data: data, width: width, height: height}
}

func (g *ArrayGrid[T]) Get(point geo2d.Point) (value T, found bool) {
	if !g.InBounds(point) {
		return lo.Empty[T](), false
	}
	return g.data[g.width*point.Y+point.X].value, true
}

func (g *ArrayGrid[T]) Set(point geo2d.Point, value T) bool {
	if !g.InBounds(point) {
		return false
	}
	g.data[g.width*point.Y+point.X] = item[T]{value, true}
	return true
}

func (g *ArrayGrid[T]) Clear(point geo2d.Point) bool {
	if !g.InBounds(point) {
		return false
	}
	g.data[g.width*point.Y+point.X] = item[T]{lo.Empty[T](), false}
	return true
}

func (g *ArrayGrid[T]) MinX() int {
	return 0
}

func (g *ArrayGrid[T]) MinY() int {
	return 0
}

func (g *ArrayGrid[T]) MaxX() int {
	return g.width - 1
}

func (g *ArrayGrid[T]) MaxY() int {
	return g.height - 1
}

func (g *ArrayGrid[T]) Bounds() (minX int, minY int, maxX int, maxY int) {
	return g.MinX(), g.MinY(), g.MaxX(), g.MaxY()
}

func (g *ArrayGrid[T]) InBounds(point geo2d.Point) bool {
	return point.X >= g.MinX() && point.X <= g.MaxX() && point.Y >= g.MinY() && point.Y <= g.MaxY()
}

func (g *ArrayGrid[T]) Points() <-chan geo2d.Point {
	c := make(chan geo2d.Point)
	go func() {
		defer close(c)

		for i := 0; i < g.width*g.height; i++ {
			c <- geo2d.Point{X: i % g.width, Y: i / g.width}
		}
	}()
	return c
}

func (g *ArrayGrid[T]) PointsSet() <-chan geo2d.Point {
	c := make(chan geo2d.Point)
	go func() {
		defer close(c)

		for i, dataItem := range g.data {
			if dataItem.isSet {
				c <- geo2d.Point{X: i % g.width, Y: i / g.width}
			}
		}
	}()
	return c
}

func (g *ArrayGrid[T]) DrawPointGrid(mapping map[T]rune, fallback rune) {
	for i, dataItem := range g.data {
		if i > 0 && i%g.width == 0 {
			print("\n")
		}

		if character, valueExists := mapping[dataItem.value]; dataItem.isSet && valueExists {
			fmt.Print(string(character))
		} else {
			fmt.Print(fallback)
		}
	}
}

func (g *ArrayGrid[T]) DrawPointGridBy(mapping func(value T, isSet bool, x int, y int) rune) {
	for i, dataItem := range g.data {
		if i > 0 && i%g.width == 0 {
			print("\n")
		}

		x := i % g.width
		y := i / g.width
		char := mapping(dataItem.value, dataItem.isSet, x, y)
		fmt.Print(string(char))
	}
}

var _ Grid[struct{}] = &ArrayGrid[struct{}]{}
