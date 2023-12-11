package day9

import (
	"2023/src/framework"
	"2023/src/framework/math"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

func Task1(data string) (result framework.Result[int]) {
	reports := parse(data)

	result.Value = lo.Sum(lop.Map(reports, func(report []int, index int) int {
		return findNextNumber(report)
	}))

	return
}

func Task2(data string) (result framework.Result[int]) {
	reports := parse(data)

	result.Value = lo.Sum(lop.Map(reports, func(report []int, index int) int {
		return findPreviousNumber(report)
	}))

	return
}

func parse(data string) [][]int {
	return lo.Map(framework.Lines(data), func(item string, index int) []int {
		return math.ExtractNumbers(item)
	})
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
	for index := range framework.RangeGen(0, len(input)-1, 1) {
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
