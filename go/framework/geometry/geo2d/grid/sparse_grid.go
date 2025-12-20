package grid

import (
	"aoc/framework"
	"aoc/framework/geometry/geo2d"
	"fmt"
)

type SparseGrid[T comparable] struct {
	data                   map[geo2d.Point]T
	minX, minY, maxX, maxY int
	dirty                  bool
}

func NewSparseGrid[T comparable]() *SparseGrid[T] {
	return &SparseGrid[T]{
		data: make(map[geo2d.Point]T),
	}
}

func (g *SparseGrid[T]) Get(key geo2d.Point) (T, bool) {
	value, found := g.data[key]
	return value, found
}

func (g *SparseGrid[T]) Set(key geo2d.Point, value T) bool {
	g.data[key] = value
	g.dirty = true
	return true
}

func (g *SparseGrid[T]) Clear(key geo2d.Point) bool {
	delete(g.data, key)
	g.dirty = true
	return true
}

func (g *SparseGrid[T]) MinX() int {
	if g.dirty {
		g.updateBounds()
	}
	return g.minX
}

func (g *SparseGrid[T]) MinY() int {
	if g.dirty {
		g.updateBounds()
	}
	return g.minY
}

func (g *SparseGrid[T]) MaxX() int {
	if g.dirty {
		g.updateBounds()
	}
	return g.maxX
}

func (g *SparseGrid[T]) MaxY() int {
	if g.dirty {
		g.updateBounds()
	}
	return g.maxY
}

func (g *SparseGrid[T]) Bounds() (int, int, int, int) {
	if g.dirty {
		g.updateBounds()
	}
	return g.minX, g.minY, g.maxX, g.maxY
}

func (g *SparseGrid[T]) InBounds(point geo2d.Point) bool {
	if g.dirty {
		g.updateBounds()
	}
	return point.X >= g.minX && point.X <= g.maxX && point.Y >= g.minY && point.Y <= g.maxY
}

func (g *SparseGrid[T]) Points() <-chan geo2d.Point {
	if g.dirty {
		g.updateBounds()
	}

	c := make(chan geo2d.Point)
	go func() {
		defer close(c)

		width := g.width()
		height := g.height()

		for i := g.minX; i < width*height; i++ {
			c <- geo2d.Point{X: i % width, Y: i / width}
		}
	}()
	return c
}

func (g *SparseGrid[T]) PointsSet() <-chan geo2d.Point {
	c := make(chan geo2d.Point)
	go func() {
		defer close(c)

		for point := range g.data {
			c <- point
		}
	}()
	return c
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
	for y := range framework.RangeGen(g.minY, g.height(), 1) {
		for x := range framework.RangeGen(g.minX, g.width(), 1) {
			value, found := g.data[geo2d.Point{X: x, Y: y}]

			if mapping != nil {
				fmt.Print(string(mapping(value, found, x, y)))
				continue
			}
		}
		print("\n")
	}
}

func (g *SparseGrid[T]) width() int {
	return g.maxX - g.minX + 1
}

func (g *SparseGrid[T]) height() int {
	return g.maxY - g.minY + 1
}

func (g *SparseGrid[T]) updateBounds() {
	g.minX = 0
	g.minY = 0
	g.maxX = 0
	g.maxY = 0

	if len(g.data) == 0 {
		return
	}

	i := 0
	for point, _ := range g.data {
		if i == 0 || point.X < g.minX {
			g.minX = point.X
		}
		if i == 0 || point.Y < g.minY {
			g.minY = point.Y
		}
		if i == 0 || point.X > g.maxX {
			g.maxX = point.X
		}
		if i == 0 || point.Y > g.maxY {
			g.maxY = point.Y
		}
		i++
	}

	g.dirty = false
}

var _ Grid[struct{}] = &SparseGrid[struct{}]{}
