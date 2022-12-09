package day8

import (
	"2022/src/framework"
	"2022/src/framework/test"
	"strconv"
)

func Task1(filePath string) (result test.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	grid := loadGrid(data)
	mask := determineVisibility(grid)

	for x, items := range mask {
		for y := range items {
			if mask[x][y] {
				result.Value++
			}
		}
	}

	return
}

func Task2(filePath string) (result test.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	grid := loadGrid(data)
	scores := createMask(grid, 0)

	for x, items := range grid {
		for y := range items {
			scores[x][y] = determineScore(x, y, grid)
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

func determineVisibility(grid [][]int) [][]bool {
	mask := createMask(grid, false)

	increasingRange := framework.Range(0, len(grid), 1)
	decreasingRange := framework.Range(len(grid)-1, len(grid), -1)

	// Top to bottom
	for _, x := range increasingRange {
		max := -1
		for _, y := range increasingRange {
			if grid[x][y] > max {
				mask[x][y] = true
				max = grid[x][y]
			}
		}
	}

	// Bottom to top
	for _, x := range increasingRange {
		max := -1
		for _, y := range decreasingRange {
			if grid[x][y] > max {
				mask[x][y] = true
				max = grid[x][y]
			}
		}
	}

	// Left to right
	for _, y := range increasingRange {
		max := -1
		for _, x := range increasingRange {
			if grid[x][y] > max {
				mask[x][y] = true
				max = grid[x][y]
			}
		}
	}

	// Right to left
	for _, y := range increasingRange {
		max := -1
		for _, x := range decreasingRange {
			if grid[x][y] > max {
				mask[x][y] = true
				max = grid[x][y]
			}
		}
	}

	return mask
}

func determineScore(x int, y int, grid [][]int) int {
	leftRange := framework.Range(x-1, x, -1)
	rightRange := framework.Range(x+1, len(grid)-x-1, 1)
	topRange := framework.Range(y-1, y, -1)
	bottomRange := framework.Range(y+1, len(grid[x])-y-1, 1)

	var scoreLeft, scoreRight, scoreTop, scoreBottom int

	// To top
	for _, newY := range topRange {
		scoreTop++
		if grid[x][newY] >= grid[x][y] {
			break
		}
	}

	// To bottom
	for _, newY := range bottomRange {
		scoreBottom++
		if grid[x][newY] >= grid[x][y] {
			break
		}
	}

	// To left
	for _, newX := range leftRange {
		scoreLeft++
		if grid[newX][y] >= grid[x][y] {
			break
		}
	}

	// To right
	for _, newX := range rightRange {
		scoreRight++
		if grid[newX][y] >= grid[x][y] {
			break
		}
	}

	return scoreLeft * scoreRight * scoreTop * scoreBottom
}

func loadGrid(data []string) [][]int {
	var grid [][]int
	for _, row := range data {
		for x, item := range []rune(row) {
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

func createMask[I any, O any](data [][]I, initial O) [][]O {
	sizeA := len(data)
	mask := make([][]O, sizeA)
	for a, aValues := range data {
		sizeB := len(data)
		aNew := make([]O, sizeB)
		for b := range aValues {
			aNew[b] = initial
		}
		mask[a] = aNew
	}
	return mask
}
