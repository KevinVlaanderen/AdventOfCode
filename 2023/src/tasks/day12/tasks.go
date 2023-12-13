package day12

import (
	"2023/src/framework"
	"2023/src/framework/math"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"strings"
)

func Task1(data string) (result framework.Result[int]) {
	lines := framework.Lines(data)

	for _, line := range lines {
		recordData := strings.Split(line, " ")
		conditions := recordData[0]
		sizes := math.ExtractNumbers(recordData[1])

		cache := framework.NewSafeCache[framework.Hash64, int]()

		matches := countFits(conditions, sizes, cache)
		result.Value += matches
	}

	return
}

func Task2(data string) (result framework.Result[int]) {
	lines := framework.Lines(data)

	cache := framework.NewSafeCache[framework.Hash64, int]()

	lop.ForEach(lines, func(line string, index int) {
		recordData := strings.Split(line, " ")
		conditions := framework.RepeatString(recordData[0], 5, "?")
		sizes := framework.RepeatSlice(math.ExtractNumbers(recordData[1]), 5)

		matches := countFits(conditions, sizes, cache)
		result.Value += matches
	})

	return
}

func countFits(conditions string, sizes []int, cache *framework.SafeCache[framework.Hash64, int]) int {
	numConditions := len(conditions)
	totalDashes := lo.Count([]rune(conditions), '#')

	var innerCountFits func(conditionIndex int, sizeIndex int, dashesFound int) int
	innerCountFits = func(conditionIndex int, sizeIndex int, dashesFound int) int {
		if matches, ok := cache.Get(framework.ComputeHash64(HashItem{conditions[conditionIndex:], sizes[sizeIndex:]})); ok {
			return matches
		}

		isLast := sizeIndex == len(sizes)-1

		currentSize := sizes[sizeIndex]

		requiredSpace := lo.Sum(sizes[sizeIndex:]) + len(sizes) - sizeIndex - 1
		if requiredSpace > numConditions-conditionIndex {
			return 0
		}

		matches := 0

	ConditionLoop:
		for i := conditionIndex; i <= numConditions-requiredSpace; i++ {
			currentDashesFound := dashesFound

			if conditions[i] == '.' {
				continue
			}

			for j := conditionIndex; j < i; j++ {
				if conditions[j] == '#' {
					break ConditionLoop
				}
			}

			for j := i; j < i+currentSize; j++ {
				if conditions[j] == '.' {
					continue ConditionLoop
				} else if conditions[j] == '#' {
					currentDashesFound++
				}
			}

			if i+currentSize < numConditions && conditions[i+currentSize] == '#' {
				continue
			}

			if isLast {
				if currentDashesFound < totalDashes {
					continue
				}
				matches++
			} else {
				matches += innerCountFits(i+currentSize+1, sizeIndex+1, currentDashesFound)
			}
		}
		cache.Set(framework.ComputeHash64(HashItem{conditions[conditionIndex:], sizes[sizeIndex:]}), matches)
		return matches
	}

	return innerCountFits(0, 0, 0)
}

type HashItem struct {
	conditions string
	sizes      []int
}
