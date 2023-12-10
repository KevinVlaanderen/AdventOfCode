package geometry

import (
	"2023/src/framework/generators"
	"fmt"
	"github.com/samber/lo"
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

func (g *Grid[T]) Keys() []Point {
	return lo.Keys(g.points)
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

func (g *Grid[T]) Delete(key Point) {
	delete(g.points, key)
}

func (g *Grid[T]) Neighbours(point Point, mustExist bool) (neighbours []Point) {
	for x := point.X - 1; x <= point.X+1; x++ {
		for y := point.Y - 1; y <= point.Y+1; y++ {
			if x < g.xMin || x > g.xMax || y < g.yMin || y > g.yMax || (x == point.X && y == point.Y) {
				continue
			}
			neighbour := Point{x, y}
			if mustExist {
				if _, ok := g.points[neighbour]; !ok {
					continue
				}
			}
			neighbours = append(neighbours, neighbour)
		}
	}
	return
}

func (g *Grid[T]) OrthogonalNeighbours(point Point, mustExist bool) (neighbours []Point) {
	top := Point{point.X, point.Y - 1}
	if _, ok := g.points[top]; (ok || !mustExist) && g.inBounds(top) {
		neighbours = append(neighbours, top)
	}
	right := Point{point.X + 1, point.Y}
	if _, ok := g.points[right]; (ok || !mustExist) && g.inBounds(right) {
		neighbours = append(neighbours, right)
	}
	bottom := Point{point.X, point.Y + 1}
	if _, ok := g.points[bottom]; (ok || !mustExist) && g.inBounds(bottom) {
		neighbours = append(neighbours, bottom)
	}
	left := Point{point.X - 1, point.Y}
	if _, ok := g.points[left]; (ok || !mustExist) && g.inBounds(left) {
		neighbours = append(neighbours, left)
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

func (g *Grid[T]) inBounds(point Point) bool {
	return point.X >= g.xMin && point.X <= g.xMax && point.Y >= g.yMin && point.Y <= g.yMax
}
