package day1

import (
	"2022/src/framework"
	"sort"
	"strconv"
)

func Task1(filePath string) (result framework.TaskResult) {
	data, err := framework.ReadLineBlocks(filePath)
	if err != nil {
		return framework.TaskResult{Error: err}
	}

	var sums []int

	for _, block := range data {
		sum := 0

		for _, line := range block {
			if value, err := strconv.Atoi(line); err != nil {
				return framework.TaskResult{Error: err}
			} else {
				sum += value
			}
		}

		sums = append(sums, sum)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	return framework.TaskResult{Value: sums[0]}
}

func Task2(filePath string) (result framework.TaskResult) {
	data, err := framework.ReadLineBlocks(filePath)
	if err != nil {
		return framework.TaskResult{Error: err}
	}

	var sums []int

	for _, block := range data {
		sum := 0

		for _, line := range block {
			if value, err := strconv.Atoi(line); err != nil {
				return framework.TaskResult{Error: err}
			} else {
				sum += value
			}
		}

		sums = append(sums, sum)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	return framework.TaskResult{Value: framework.Sum(sums[0:3])}
}
