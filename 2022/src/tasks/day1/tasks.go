package day1

import (
	"2022/src/framework/math"
	"2022/src/framework/tasks"
	"sort"
	"strconv"
)

func Task1(filePath string) (result tasks.TaskResult[int]) {
	elfs := tasks.ReadStream(filePath, createParser())

	var sums []int
	for elf := range elfs {
		sums = append(sums, math.Sum(elf))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	return tasks.TaskResult[int]{Value: sums[0]}
}

func Task2(filePath string) (result tasks.TaskResult[int]) {
	elfs := tasks.ReadStream(filePath, createParser())

	var sums []int
	for elf := range elfs {
		sums = append(sums, math.Sum(elf))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	return tasks.TaskResult[int]{Value: math.Sum(sums[0:3])}
}

type Elf = []Calories
type Calories = int

func createParser() func(line string) (result Elf, hasResult bool, err error) {
	var elf Elf

	return func(line string) (result Elf, hasResult bool, err error) {
		if line == "" {
			result = elf
			elf = nil
			return result, true, nil
		} else if value, err := strconv.Atoi(line); err == nil {
			elf = append(elf, value)
			return result, false, nil
		} else {
			return result, false, err
		}
	}
}
