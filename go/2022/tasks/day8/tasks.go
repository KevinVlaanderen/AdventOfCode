//go:generate go run aoc/cmd/generate-tests

package day8

import (
	"aoc/2022/tasks/day8/model"
	"aoc/framework/geometry/geo2d"
	"aoc/framework/tasks"
	"go/types"
	"strconv"
)

// Task1 type:mock 	file:data	expected:21
// Task1 type:real 	file:day8 	expected:1792
func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	grid := parse(data)

	mask := grid.DetermineVisibility()

	for x, items := range mask {
		for y := range items {
			if mask[x][y] {
				result.Value++
			}
		}
	}

	return
}

// Task2 type:mock 	file:data	expected:8
// Task2 type:real 	file:day8 	expected:334880
func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	grid := parse(data)

	scores := geo2d.CreateMask(grid, 0)

	for x, items := range grid {
		for y := range items {
			scores[x][y] = grid.DetermineScore(x, y)
		}
	}

	for x, items := range scores {
		for y := range items {
			if scores[x][y] > result.Value {
				result.Value = scores[x][y]
			}
		}
	}

	return
}

func parse(data string) model.Grid {
	var grid model.Grid

	for _, line := range tasks.Lines(data) {
		for x, item := range []rune(line) {
			value, _ := strconv.Atoi(string(item))
			if x >= len(grid) {
				grid = append(grid, []int{value})
			} else {
				grid[x] = append(grid[x], value)
			}
		}
	}

	return grid
}
