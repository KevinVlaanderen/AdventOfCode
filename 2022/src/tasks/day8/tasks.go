package day8

import (
	"2022/src/framework/geometry"
	"2022/src/framework/tasks"
	"2022/src/tasks/day8/model"
	"strconv"
)

func Task1(filePath string) (result tasks.TaskResult[int]) {
	grid := <- tasks.ReadStream(filePath, createParser())

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

func Task2(filePath string) (result tasks.TaskResult[int]) {
	grid := <- tasks.ReadStream(filePath, createParser())

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

func createParser() func(line string) (result model.Grid, hasResult bool, err error) {
	var grid model.Grid

	return func(line string) (result model.Grid, hasResult bool, err error) {
		if line == "" {
			return grid, true, nil
		}

		for x, item := range []rune(line) {
			value, _ := strconv.Atoi(string(item))
			if x >= len(grid) {
				grid = append(grid, []int{value})
			} else {
				grid[x] = append(grid[x], value)
			}
		}

		return result, false, nil
	}
}
