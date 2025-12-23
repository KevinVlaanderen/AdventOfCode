//go:generate go run aoc/cmd/generate-tests

package day1

import (
	"aoc/framework/math"
	"aoc/framework/tasks"
	"go/types"
	"sort"
	"strconv"

	"github.com/samber/lo"
)

type Elf = []Calories
type Calories = int

// Task1 type:mock 	file:data	expected:24000
// Task1 type:real 	file:day1 	expected:67027
func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	elfs := parse(data)

	var sums []int
	for _, elf := range elfs {
		sums = append(sums, math.Sum(elf))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	result.Value = sums[0]

	return
}

// Task2 type:mock 	file:data	expected:45000
// Task2 type:real 	file:day1 	expected:197291
func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
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
	return lo.Map(tasks.LineBlocks(data), func(block []string, index int) Elf {
		return lo.Map(block, func(line string, index int) Calories {
			value, _ := strconv.Atoi(line)
			return value
		})
	})
}
