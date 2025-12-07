package day4

import (
	"2025/src/framework"
	"2025/src/framework/geometry"
	"2025/src/framework/geometry/grid"
	"go/types"
)

func Task(data string, _ types.Nil) (result framework.Result[int]) {
	rolls := parse(data)
	x1, x2, y1, y2 := rolls.Boundaries()

	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			point := geometry.Point{X: x, Y: y}

			hasRoll, found := rolls.Get(&point)
			if !found || !hasRoll {
				continue
			}

			count := 0
			for _, neighbour := range point.Neighbors(geometry.All) {
				hasRoll, found := rolls.Get(&neighbour)
				if found && hasRoll {
					count++
				}
			}
			if count < 4 {
				result.Value++
			}
		}
	}

	return
}

func parse(data string) grid.Grid[bool] {
	lines := framework.CharLines(data)
	width := len(lines[0])
	height := len(lines)
	g := grid.NewGrid[bool](width, height)
	for y, line := range lines {
		for x, char := range line {
			_ = g.Set(&geometry.Point{X: x, Y: y}, char == '@')
		}
	}
	return g
}
