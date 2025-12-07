package math

import (
	"github.com/samber/lo"
	"regexp"
	"strconv"
)

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func MaxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func MinInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func AbsInt(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func Length(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

var numberPattern = regexp.MustCompile(`-?\d+`)

func ExtractNumbers(line string) []int {
	return lo.Map(numberPattern.FindAllString(line, -1), func(item string, index int) int {
		if number, err := strconv.Atoi(item); err != nil {
			panic(err)
		} else {
			return number
		}
	})
}
