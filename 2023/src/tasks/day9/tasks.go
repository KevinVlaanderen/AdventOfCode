package day9

import (
	"2023/src/framework/generators"
	"2023/src/framework/number"
	"2023/src/framework/task"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

func Task1(filePath string) (result task.Result[int]) {
	reports := task.Read(filePath, func(line string) (result []int, hasResult bool, err error) {
		if line == "" {
			return
		}
		return number.ExtractNumbers(line), true, nil
	})

	result.Value = lo.Sum(lop.Map(reports, func(report []int, index int) int {
		return findNextNumber(report)
	}))

	return
}

func findNextNumber(input []int) int {
	lastNumbers := []int{input[len(input)-1]}

	for {
		input = findDiff(input)
		if lo.EveryBy(input, func(item int) bool {
			return item == 0
		}) {
			break
		}
		lastNumbers = append(lastNumbers, input[len(input)-1])
	}

	offset := 0
	for i := len(lastNumbers) - 1; i >= 0; i-- {
		offset = lastNumbers[i] + offset
	}
	return offset
}

func findDiff(input []int) (diff []int) {
	for index := range generators.RangeGen(0, len(input)-1, 1) {
		diff = append(diff, input[index+1]-input[index])
	}
	return
}
