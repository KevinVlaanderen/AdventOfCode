//go:generate go run aoc/cmd/generate-tests

package day2

import (
	"aoc/framework/datastructures"
	math2 "aoc/framework/math"
	"aoc/framework/tasks"
	"go/types"
	"reflect"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Range struct {
	Min, Max int
}

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	pairs := parse(data)

	result.Value = lo.Reduce(pairs, func(result int, r Range, index int) int {
		return result + sumRepeats2(r)
	}, 0)

	return
}

func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	pairs := parse(data)

	result.Value = lo.Reduce(pairs, func(result int, r Range, index int) int {
		return result + sumRepeatsFactors(r)
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

func sumRepeats2(r Range) int {
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

func sumRepeatsFactors(r Range) int {
	sum := 0

numbers:
	for i := r.Min; i <= r.Max; i++ {
		factors := math2.Factors(math2.Length(i))
		digits := math2.Digits(i)

	parts:
		for _, partSize := range factors[:len(factors)-1] {
			slices := datastructures.Partition(digits, partSize)
			numSlices := len(slices)

			for sliceIndex := 1; sliceIndex < numSlices; sliceIndex++ {
				if !reflect.DeepEqual(slices[sliceIndex], slices[sliceIndex-1]) {
					continue parts
				}
			}

			sum += i
			continue numbers
		}
	}

	return sum
}
