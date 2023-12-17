package grid

import (
	"2023/src/framework/geometry"
	"errors"
	"fmt"
	"github.com/samber/lo"
)

type Grid[T comparable] struct {
	data          [][]T
	width, height int
}

func NewGrid[T comparable](width, height int) Grid[T] {
	data := make([][]T, height)
	for y := 0; y < height; y++ {
		data[y] = make([]T, width)
	}
	return Grid[T]{data: data, width: width, height: height}
}

func (g *Grid[T]) InBounds(point *geometry.Point) bool {
	return point.X >= 0 && point.X < g.width && point.Y >= 0 && point.Y < g.height
}

func (g *Grid[T]) Get(point *geometry.Point) (value T, found bool) {
	if !g.InBounds(point) {
		return lo.Empty[T](), false
	}
	return g.data[point.Y][point.X], true
}

func (g *Grid[T]) Set(point *geometry.Point, value T) error {
	if !g.InBounds(point) {
		return errors.New("out of bounds")
	}
	g.data[point.Y][point.X] = value
	return nil
}

func (g *Grid[T]) Boundaries() (int, int, int, int) {
	return 0, g.width - 1, 0, g.height - 1
}

func (g *Grid[T]) DrawPointGrid(fn func(value *T, x int, y int) (rune, bool), fallback map[T]rune) {
	for y, row := range g.data {
		for x, v := range row {
			if char, ok := fn(&v, x, y); ok {
				fmt.Print(string(char))
				continue
			}

			if character, valueExists := fallback[v]; valueExists {
				fmt.Print(string(character))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
