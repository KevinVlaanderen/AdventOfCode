package day8

import (
	"2022/src/framework"
	"strconv"
)

func Task1(filePath string) (result framework.TaskResult[int]) {
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

func Task2(filePath string) (result framework.TaskResult[int]) {
	_, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	return
}

func determineVisibility(grid [][]int) [][]bool {
	mask := createMask(grid)

	increasingRange := NewSlice(0, len(grid), 1)
	decreasingRange := NewSlice(len(grid)-1, len(grid), -1)

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

func createMask[T any](data [][]T) [][]bool {
	sizeA := len(data)
	mask := make([][]bool, sizeA)
	for a, aValues := range data {
		sizeB := len(data)
		aNew := make([]bool, sizeB)
		for b, _ := range aValues {
			aNew[b] = false
		}
		mask[a] = aNew
	}
	return mask
}

func NewSlice(start, count, step int) []int {
	s := make([]int, count)
	for i := range s {
		s[i] = start
		start += step
	}
	return s
}
