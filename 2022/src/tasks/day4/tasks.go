package day4

import (
	"2022/src/framework"
	"fmt"
	"strconv"
	"strings"
)

func Task1(filePath string) (result framework.TaskResult) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		return framework.TaskResult{Error: err}
	}

	for _, line := range data {
		parts := strings.Split(line, ",")
		var part1Low, part1High, part2Low, part2High int
		if part1Low, part1High, err = GetRange(parts[0]); err != nil {
			return framework.TaskResult{Error: fmt.Errorf("failed to get range for part 1: %v", err)}
		}
		if part2Low, part2High, err = GetRange(parts[1]); err != nil {
			return framework.TaskResult{Error: fmt.Errorf("failed to get range for part 2: %v", err)}
		}

		if (part1Low <= part2Low && part1High >= part2High) || (part2Low <= part1Low && part2High >= part1High) {
			result.Value++
		}
	}
	return
}

func Task2(filePath string) (result framework.TaskResult) {
	_, err := framework.ReadLines(filePath)
	if err != nil {
		return framework.TaskResult{Error: err}
	}

	return
}

func GetRange(input string) (low int, high int, err error) {
	parts := strings.Split(input, "-")
	if low, err = strconv.Atoi(parts[0]); err != nil {
		return
	}
	if high, err = strconv.Atoi(parts[1]); err != nil {
		return
	}

	return
}
