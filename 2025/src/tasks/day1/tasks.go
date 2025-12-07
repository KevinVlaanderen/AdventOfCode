package day1

import (
	"2025/src/framework"
	"go/types"
	"strconv"
)

func Task1(data string, _ types.Nil) (result framework.Result[int]) {
	position := 50

	for _, line := range framework.Lines(data) {
		amount, _ := strconv.Atoi(line[1:])

		if line[0] == 'L' {
			position -= amount % 100
		} else {
			position += amount % 100
		}

		if position < 0 {
			position = 100 + position
		} else if position > 99 {
			position = position % 100
		}

		if position == 0 {
			result.Value++
		}
	}

	return
}
