package day12

import (
	"2023/src/framework"
	"2023/src/framework/math"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"strings"
)

func Task1(data string) (result framework.Result[int]) {
	lines := framework.Lines(data)

	for _, line := range lines {
		recordData := strings.Split(line, " ")
		conditions := recordData[0]
		sizes := math.ExtractNumbers(recordData[1])

		matches := countFits(conditions, sizes, 0, 0)
		result.Value += matches
	}

	return
}

func countFits(conditions string, sizes []int, conditionIndex int, sizeIndex int) int {
	isLast := sizeIndex == len(sizes)-1
	currentSize := sizes[sizeIndex]

	requiredSpace := lo.Sum(sizes[sizeIndex:]) + len(sizes[sizeIndex:]) - 1
	if requiredSpace > len(conditions[conditionIndex:]) {
		return 0
	}

	matches := 0

	for i := conditionIndex; i <= len(conditions)-requiredSpace; i++ {
		if conditions[i] == '.' {
			continue
		}

		if slices.Contains([]rune(conditions[i:i+currentSize]), '.') {
			continue
		} else if slices.Contains([]rune(conditions[conditionIndex:i]), '#') {
			break
		} else if i+currentSize < len(conditions) && conditions[i+currentSize] == '#' {
			continue
		}

		nextIndex := i + currentSize + 1
		remaining := conditions[nextIndex-1:]

		if isLast {
			if strings.ContainsRune(remaining, '#') {
				continue
			} else {
				matches++
			}
		} else {
			matches += countFits(conditions, sizes, nextIndex, sizeIndex+1)
		}
	}
	return matches
}

//func countFits(groupConditions []GroupCondition, groupSizes []int) int {
//	count := 0
//
//	return count
//}
//
//type Setup struct {
//	groupConditions []GroupCondition
//	groupSizes      []int
//}
//
//func NewRecord(data string) Setup {
//	parts := strings.Split(data, " ")
//	groupConditionsData := parts[0]
//	groupSizesData := parts[1]
//
//	return Setup{groupConditions: parseGroupConditions(groupConditionsData), groupSizes: parseGroupSizes(groupSizesData)}
//}
//
//func parseGroupConditions(data string) []GroupCondition {
//	groupConditions := make([]GroupCondition, 0)
//
//	currentGroupCondition := GroupCondition{}
//	for _, char := range data {
//		switch char {
//		case '.':
//			if len(currentGroupCondition) > 0 {
//				groupConditions = append(groupConditions, currentGroupCondition)
//				currentGroupCondition = GroupCondition{}
//			}
//		case '#':
//			currentGroupCondition = append(currentGroupCondition, Operational)
//		case '?':
//			currentGroupCondition = append(currentGroupCondition, Unknown)
//		}
//	}
//	if len(currentGroupCondition) > 0 {
//		groupConditions = append(groupConditions, currentGroupCondition)
//		currentGroupCondition = GroupCondition{}
//	}
//	return groupConditions
//}
//
//func parseGroupSizes(data string) []int {
//	groupSizes := make([]int, len(data)/2+1)
//
//	lo.ForEach(strings.Split(data, ","), func(item string, index int) {
//		if n, err := strconv.Atoi(item); err != nil {
//			panic(err)
//		} else {
//			groupSizes = append(groupSizes, n)
//		}
//	})
//	return groupSizes
//}
//
//type GroupCondition []SpringCondition
//type SpringCondition uint8
//
//const (
//	Operational SpringCondition = iota
//	Unknown
//	Damaged
//)
