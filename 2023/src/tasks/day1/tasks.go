package day1

import (
	"2023/src/framework"
	"2023/src/framework/parse"
	"fmt"
	"github.com/samber/lo"
	"golang.org/x/exp/maps"
	"strconv"
	"strings"
)

var digits = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}
var words = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func Task1(filePath string) (result framework.Result[int]) {
	lookup := digits

	for line := range framework.ReadStream(filePath, parse.Lines()) {
		result.Value += calculateCalibrationValue(line, lookup)
	}
	return
}

func Task2(filePath string) (result framework.Result[int]) {
	lookup := lo.Assign(digits, words)

	for line := range framework.ReadStream(filePath, parse.Lines()) {
		result.Value += calculateCalibrationValue(line, lookup)
	}
	return
}

func calculateCalibrationValue(line string, lookup map[string]int) int {
	var firstDigit, lastDigit int
	var firstIndex, lastIndex int

	for _, key := range maps.Keys(lookup) {
		index := strings.Index(line, key)
		if index >= 0 && (firstDigit == 0 || index < firstIndex) {
			firstDigit = lookup[key]
			firstIndex = index
		}

		index = strings.LastIndex(line, key)
		if index >= 0 && (lastDigit == 0 || index > lastIndex) {
			lastDigit = lookup[key]
			lastIndex = index
		}
	}

	if value, err := strconv.Atoi(fmt.Sprintf("%v%v", firstDigit, lastDigit)); err != nil {
		panic(err)
	} else {
		return value
	}
}
