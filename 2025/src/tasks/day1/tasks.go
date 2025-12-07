package day1

import (
	"2025/src/framework"
	"2025/src/framework/math"
	"strconv"
)

func Task(data string, countPassing bool) (result framework.Result[int]) {
	position := 50
	ending, passing := 0, 0

	for _, line := range framework.Lines(data) {
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
