//go:generate go run aoc/cmd/generate-tests

package day8

import (
	"aoc/2023/tasks/day8/model"
	"aoc/framework/math"
	"aoc/framework/tasks"
	"go/types"

	lop "github.com/samber/lo/parallel"
)

// Task1 type:mock 	file:data	expected:2
// Task1 type:mock 	file:data2	expected:6
// Task1 type:real 	file:day8 	expected:14893
func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	blocks := tasks.Blocks(data)
	network := model.NewNetwork(blocks[0], blocks[1])

	start := network.IndexOf("AAA")

	result.Value = network.FindRequiredStepsBy(start, func(item int) bool {
		return network.NameOf(item) == "ZZZ"
	})

	return
}

// Task2 type:mock 	file:data3	expected:6
// Task2 type:real 	file:day8 	expected:10241191004509
func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	blocks := tasks.Blocks(data)
	network := model.NewNetwork(blocks[0], blocks[1])

	startList := network.IndicesBy(func(item string, index int) bool {
		return item[2] == 'A'
	})

	requiredSteps := lop.Map(startList, func(start int, index int) int {
		return network.FindRequiredStepsBy(start, func(item int) bool {
			return network.NameOf(item)[2] == 'Z'
		})
	})

	result.Value = math.LCM(requiredSteps[0], requiredSteps[1], requiredSteps[2:]...)

	return
}
