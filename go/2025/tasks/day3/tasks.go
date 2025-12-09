package day3

import (
	"aoc/framework/tasks"
	"strconv"

	"github.com/samber/lo"
)

type Battery struct {
	Value, Position int
}

func Task(data string, amount int) (result tasks.Result[int]) {
	banks := parse(data)

	for _, bank := range banks {
		values := make([]int, amount)
		start := 0

		for i := 0; i < amount; i++ {
			value, position := findHighest(bank, start, len(bank)-amount+i+1)
			values[i] = value
			start = position + 1
		}

		strValue := lo.Reduce(values, func(agg string, item int, index int) string {
			return agg + strconv.Itoa(item)
		}, "")

		value, _ := strconv.Atoi(strValue)
		result.Value += value
	}

	return
}

func parse(data string) [][]int {
	return lo.Map(tasks.CharLines(data), func(line []rune, index int) []int {
		return lo.Map(line, func(char rune, index int) int {
			return int(char) - 48
		})
	})
}

func findHighest(bank []int, start, end int) (value, position int) {
	for i := start; i < end; i++ {
		if bank[i] > value {
			value = bank[i]
			position = i
		}
	}
	return
}
