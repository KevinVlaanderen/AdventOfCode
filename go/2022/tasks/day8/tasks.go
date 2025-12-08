package day8

import (
	"aoc/2022/tasks/day8/model"
	"aoc/framework"
	"aoc/framework/geometry"
	"go/types"
	"strconv"
)

func Task1(data string, _ types.Nil) (result framework.Result[int]) {
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

func Task2(data string, _ types.Nil) (result framework.Result[int]) {
	grid := parse(data)

	scores := geometry.CreateMask(grid, 0)

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

	for _, line := range framework.Lines(data) {
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
