package day15

import (
	"2023/src/framework"
	"strings"
)

func Task1(data string) (result framework.Result[int]) {
	chars := framework.Lines(data)[0]

	var current int
	var text strings.Builder
	for _, char := range chars {
		if char == ',' {
			result.Value += current
			current = 0
			text.Reset()
			continue
		}
		current = ((current + int(char)) * 17) % 256
	}
	result.Value += current

	return
}
