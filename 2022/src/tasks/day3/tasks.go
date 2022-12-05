package day3

import (
	"2022/src/framework"
	"fmt"
	"unicode"
)

func Task1(filePath string) (result framework.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	for _, line := range data {
		runes := []rune(line)
		length := len(runes)
		compartment1 := runes[:length/2]
		compartment2 := runes[length/2:]

		if difference, err := framework.Intersection(compartment1, compartment2); err != nil {
			result.Error = err
			return
		} else {
			if len(difference) != 1 {
				result.Error = fmt.Errorf("invalid number of differences in %s", line)
				return
			}

			result.Value += GetPriority(difference[0])
		}
	}

	return
}

func Task2(filePath string) (result framework.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	for i := 0; i < len(data)/3; i++ {
		group := data[i*3 : i*3+3]
		if difference, err := framework.Intersection([]rune(group[0]), []rune(group[1]), []rune(group[2])); err != nil {
			result.Error = err
			return
		} else {
			if len(difference) != 1 {
				result.Error = fmt.Errorf("invalid number of differences in group %s", group)
				return
			}

			result.Value += GetPriority(difference[0])
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
