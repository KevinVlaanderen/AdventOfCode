//go:generate go run aoc/cmd/generate-tests

package day4

import (
	"aoc/framework/geometry/geo2d"
	"aoc/framework/geometry/geo2d/grid"
	"aoc/framework/tasks"
	"go/types"
)

// Task1 type:mock 	file:data	expected:13
// Task1 type:real 	file:day4 	expected:1491
func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	rolls := parse(data)
	minX, minY, maxX, maxY := rolls.Bounds()

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if canAccess(rolls, geo2d.Point{X: x, Y: y}) {
				result.Value++
			}
		}
	}

	return
}

// Task2 type:mock 	file:data	expected:43
// Task2 type:real 	file:day4 	expected:8722
func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	rolls := parse(data)
	minX, minY, maxX, maxY := rolls.Bounds()
	removed := 0
	keepTrying := true

	for keepTrying {
		keepTrying = false

		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				point := geo2d.Point{X: x, Y: y}
				if canAccess(rolls, point) {
					_ = rolls.Set(point, false)
					removed++
					keepTrying = true
				}
			}
		}
	}

	result.Value = removed

	return
}

func parse(data string) grid.Grid[bool] {
	lines := tasks.CharLines(data)
	width := len(lines[0])
	height := len(lines)
	g := grid.NewArrayGrid[bool](width, height)
	for y, line := range lines {
		for x, char := range line {
			_ = g.Set(geo2d.Point{X: x, Y: y}, char == '@')
		}
	}
	return g
}

func canAccess(rolls grid.Grid[bool], point geo2d.Point) bool {
	hasRoll, found := rolls.Get(point)
	if !found || !hasRoll {
		return false
	}

	count := 0
	for _, neighbour := range point.Neighbors(geo2d.All) {
		hasRoll, found = rolls.Get(neighbour)
		if found && hasRoll {
			count++
		}
	}
	return count < 4
}
