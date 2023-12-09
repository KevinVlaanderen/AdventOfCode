package number

import (
	"github.com/samber/lo"
	"regexp"
	"strconv"
)

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
