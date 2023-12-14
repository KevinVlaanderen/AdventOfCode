package day14

import (
	"2023/src/framework"
)

func Task1(data string) (result framework.Result[int]) {
	lines := framework.Lines(data)

	result.Value = calculateLoad(lines)

	return
}

func calculateLoad(lines []string) (result int) {
	height := len(lines)
	highest := framework.Range(height, len(lines[0]), 0)

	for y, line := range lines {
		for x, char := range line {
			switch char {
			case 'O':
				result += highest[x]
				highest[x] -= 1
			case '#':
				highest[x] = height - y - 1
			}
		}
	}
	return
}
