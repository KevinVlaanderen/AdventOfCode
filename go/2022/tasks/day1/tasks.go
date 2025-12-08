package day1

import (
	"aoc/framework"
	"aoc/framework/math"
	"go/types"
	"sort"
	"strconv"

	"github.com/samber/lo"
)

type Elf = []Calories
type Calories = int

func Task1(data string, _ types.Nil) (result framework.Result[int]) {
	elfs := parse(data)

	var sums []int
	for _, elf := range elfs {
		sums = append(sums, math.Sum(elf))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	result.Value = sums[0]

	return
}

func Task2(data string, _ types.Nil) (result framework.Result[int]) {
	elfs := parse(data)

	var sums []int
	for _, elf := range elfs {
		sums = append(sums, math.Sum(elf))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	result.Value = math.Sum(sums[0:3])

	return
}

func parse(data string) []Elf {
	return lo.Map(framework.LineBlocks(data), func(block []string, index int) Elf {
		return lo.Map(block, func(line string, index int) Calories {
			value, _ := strconv.Atoi(line)
			return value
		})
	})
}
