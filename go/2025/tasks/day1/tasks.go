//go:generate go run aoc/cmd/generate-tests

package day1

import (
	"aoc/framework/math"
	"aoc/framework/tasks"
	"strconv"
)

func Task1(data string, countPassing bool) (result tasks.Result[int]) {
	return solve(data, countPassing)
}

func Task2(data string, countPassing bool) (result tasks.Result[int]) {
	return solve(data, countPassing)
}

func solve(data string, countPassing bool) (result tasks.Result[int]) {
	position := 50
	ending, passing := 0, 0

	for _, line := range tasks.Lines(data) {
		rotation, _ := strconv.Atoi(line[1:])
		if line[0] == 'L' {
			rotation *= -1
		}

		position, ending, passing = countZeroes(position, rotation)
		if countPassing {
			result.Value += ending + passing
		} else {
			result.Value += ending
		}
	}

	return
}

func countZeroes(position, rotation int) (newPosition, ending, passing int) {
	normalizedRotation := rotation % 100
	fullRotations := math.AbsInt(rotation) / 100

	passing += fullRotations

	if math.AbsInt(normalizedRotation) == 0 {
		return position, ending, passing
	}

	position += normalizedRotation

	if position == 0 || position == 100 {
		ending++
	} else if (position < 0 && position != normalizedRotation) || position > 100 {
		passing++
	}

	if position >= 100 {
		position -= 100
	} else if position < 0 {
		position += 100
	}

	return position, ending, passing
}
