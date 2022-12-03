package tasks

import (
	"2022/src/framework"
	"sort"
	"strconv"
)

type Day1 struct{}

func (d Day1) Task1(filePath string) (result int, err error) {
	data, err := framework.ReadLineBlocks(filePath)

	var sums []int

	for _, block := range data {
		sum := 0

		for _, line := range block {
			if value, err := strconv.Atoi(line); err != nil {
				return -1, err
			} else {
				sum += value
			}
		}

		sums = append(sums, sum)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	return sums[0], nil
}

func (d Day1) Task2(filePath string) (result int, err error) {
	data, err := framework.ReadLineBlocks(filePath)

	var sums []int

	for _, block := range data {
		sum := 0

		for _, line := range block {
			if value, err := strconv.Atoi(line); err != nil {
				return -1, err
			} else {
				sum += value
			}
		}

		sums = append(sums, sum)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	result = framework.Sum(sums[0:3])

	return
}
