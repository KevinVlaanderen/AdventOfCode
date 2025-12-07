package day2

import (
	"2025/src/framework"
	math2 "2025/src/framework/math"
	"go/types"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Range struct {
	Min, Max int
}

func Task1(data string, _ types.Nil) (result framework.Result[int]) {
	pairs := parse(data)

	result.Value = lo.Reduce(pairs, func(result int, r Range, index int) int {
		return result + sumRepeats(r)
	}, 0)

	return
}

func parse(data string) []Range {
	pairs := strings.Split(data, ",")
	return lo.Map(pairs, func(pair string, index int) Range {
		minMax := strings.Split(pair, "-")
		minRange, _ := strconv.Atoi(minMax[0])
		maxRange, _ := strconv.Atoi(minMax[1])
		return Range{
			Min: minRange,
			Max: maxRange,
		}
	})
}

func sumRepeats(r Range) int {
	sum := 0

	for i := r.Min; i <= r.Max; i++ {
		numDigits, _ := math2.Length2(i)
		if numDigits%2 != 0 {
			continue
		}

		halveDigits := numDigits / 2

		highPart := i / math2.PowInt(10, halveDigits)
		lowPart := i - highPart*math2.PowInt(10, halveDigits)

		if lowPart == highPart {
			sum += i
		}
	}

	return sum
}
