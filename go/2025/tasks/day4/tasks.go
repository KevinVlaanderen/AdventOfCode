package day4

import (
	"aoc/framework/geometry/geo2d"
	"aoc/framework/geometry/geo2d/grid"
	"aoc/framework/tasks"
	"go/types"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	rolls := parse(data)
	x1, x2, y1, y2 := rolls.Boundaries()

	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			if canAccess(rolls, &geo2d.Point{X: x, Y: y}) {
				result.Value++
			}
		}
	}

	return
}

func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	rolls := parse(data)
	x1, x2, y1, y2 := rolls.Boundaries()
	removed := 0
	keepTrying := true

	for keepTrying {
		keepTrying = false

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				point := &geo2d.Point{X: x, Y: y}
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
	g := grid.NewGrid[bool](width, height)
	for y, line := range lines {
		for x, char := range line {
			_ = g.Set(&geo2d.Point{X: x, Y: y}, char == '@')
		}
	}
	return g
}

func canAccess(rolls grid.Grid[bool], point *geo2d.Point) bool {
	hasRoll, found := rolls.Get(point)
	if !found || !hasRoll {
		return false
	}

	count := 0
	for _, neighbour := range point.Neighbors(geo2d.All) {
		hasRoll, found := rolls.Get(&neighbour)
		if found && hasRoll {
			count++
		}
	}
	return count < 4
}
