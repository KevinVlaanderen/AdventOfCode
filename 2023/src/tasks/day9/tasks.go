package day9

import (
	"2023/src/framework/generators"
	"2023/src/framework/number"
	"2023/src/framework/task"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

func Task1(filePath string) (result task.Result[int]) {
	reports := task.Read(filePath, parse)

	result.Value = lo.Sum(lop.Map(reports, func(report []int, index int) int {
		return findNextNumber(report)
	}))

	return
}

func Task2(filePath string) (result task.Result[int]) {
	reports := task.Read(filePath, parse)

	result.Value = lo.Sum(lop.Map(reports, func(report []int, index int) int {
		return findPreviousNumber(report)
	}))

	return
}

func parse(line string) (result []int, hasResult bool, err error) {
	if line == "" {
		return
	}
	return number.ExtractNumbers(line), true, nil
}

func findNextNumber(input []int) int {
	lastNumbers := []int{input[len(input)-1]}
	var done bool
	for {
		if input, done = findDiff(input); done {
			break
		} else {
			lastNumbers = append(lastNumbers, input[len(input)-1])
		}
	}
	return extrapolate(lastNumbers, Next)
}

func findPreviousNumber(input []int) int {
	firstNumbers := []int{input[0]}
	var done bool
	for {
		if input, done = findDiff(input); done {
			break
		} else {
			firstNumbers = append(firstNumbers, input[0])
		}
	}
	return extrapolate(firstNumbers, Previous)
}

func extrapolate(numbers []int, direction Direction) int {
	offset := 0
	for i := len(numbers) - 1; i >= 0; i-- {
		if direction == Next {
			offset = numbers[i] + offset
		} else {
			offset = numbers[i] - offset
		}
	}
	return offset
}

func findDiff(input []int) (diff []int, done bool) {
	for index := range generators.RangeGen(0, len(input)-1, 1) {
		diff = append(diff, input[index+1]-input[index])
	}
	done = lo.EveryBy(diff, func(item int) bool {
		return item == 0
	})
	return
}

type Direction int

const (
	Next Direction = iota
	Previous
)
