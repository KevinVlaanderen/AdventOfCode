package day3

import (
	"2025/src/framework"
	"go/types"
	"strconv"

	"github.com/samber/lo"
)

func Task1(data string, _ types.Nil) (result framework.Result[int]) {
	banks := parse(data)

	for _, bank := range banks {
		firstValue, firstPosition := findHighest(bank, 0, len(bank)-1)
		secondValue, _ := findHighest(bank, firstPosition+1, len(bank))

		value, _ := strconv.Atoi(strconv.Itoa(firstValue) + strconv.Itoa(secondValue))
		result.Value += value
	}

	return
}

func parse(data string) [][]int {
	return lo.Map(framework.CharLines(data), func(line []rune, index int) []int {
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
