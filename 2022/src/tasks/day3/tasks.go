package day3

import (
	"2022/src/framework"
	"fmt"
	"math/bits"
	"unicode"
)

func Task1(filePath string) (result framework.TaskResult) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		return framework.TaskResult{Error: err}
	}

	for _, line := range data {
		runes := []rune(line)
		length := len(runes)
		compartment1 := runes[:length/2]
		compartment2 := runes[length/2:]

		if difference, err := Intersection(compartment1, compartment2); err != nil {
			return framework.TaskResult{Error: err}
		} else {
			if len(difference) != 1 {
				return framework.TaskResult{Error: fmt.Errorf("invalid number of differences in %s", line)}
			}

			result.Value += GetPriority(difference[0])
		}
	}

	return
}

func Task2(filePath string) (result framework.TaskResult) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		return framework.TaskResult{Error: err}
	}

	for i := 0; i < len(data)/3; i++ {
		group := data[i*3 : i*3+3]
		if difference, err := Intersection([]rune(group[0]), []rune(group[1]), []rune(group[2])); err != nil {
			return framework.TaskResult{Error: err}
		} else {
			if len(difference) != 1 {
				return framework.TaskResult{Error: fmt.Errorf("invalid number of differences in group %s", group)}
			}

			result.Value += GetPriority(difference[0])
		}
	}

	return
}

func Intersection[T comparable](slices ...[]T) (intersection []T, err error) {
	num := len(slices)
	if num > 64 {
		return []T{}, fmt.Errorf("too many slices provides: %d", num)
	}

	m := make(map[T]uint64)

	for index, slice := range slices {
		for _, k := range slice {
			m[k] |= 1 << index
		}
	}

	for k, v := range m {
		if bits.OnesCount64(v) == num {
			intersection = append(intersection, k)
		}
	}

	return
}

func GetPriority(item rune) int {
	if unicode.IsLower(item) {
		return int(item) - 96
	} else {
		return int(item) - 38
	}
}
