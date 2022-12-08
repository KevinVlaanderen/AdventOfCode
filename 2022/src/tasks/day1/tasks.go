package day1

import (
	"2022/src/framework"
	"2022/src/framework/test"
	"sort"
	"strconv"
)

func Task1(filePath string) (result test.TaskResult[int]) {
	data, err := framework.ReadLineBlocks(filePath)
	if err != nil {
		result.Error = err
		return
	}

	var sums []int

	for _, block := range data {
		sum := 0

		for _, line := range block {
			if value, err := strconv.Atoi(line); err != nil {
				result.Error = err
				return
			} else {
				sum += value
			}
		}

		sums = append(sums, sum)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	result.Value = sums[0]
	return
}

func Task2(filePath string) (result test.TaskResult[int]) {
	data, err := framework.ReadLineBlocks(filePath)
	if err != nil {
		result.Error = err
		return
	}

	var sums []int

	for _, block := range data {
		sum := 0

		for _, line := range block {
			if value, err := strconv.Atoi(line); err != nil {
				result.Error = err
				return
			} else {
				sum += value
			}
		}

		sums = append(sums, sum)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	result.Value = framework.Sum(sums[0:3])
	return
}
