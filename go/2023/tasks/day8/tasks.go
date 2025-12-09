package day8

import (
	"2023/src/tasks/day8/model"
	"aoc/framework/math"
	"aoc/framework/tasks"
	"go/types"

	lop "github.com/samber/lo/parallel"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	blocks := tasks.Blocks(data)
	network := model.NewNetwork(blocks[0], blocks[1])

	start := network.IndexOf("AAA")

	result.Value = network.FindRequiredStepsBy(start, func(item int) bool {
		return network.NameOf(item) == "ZZZ"
	})

	return
}

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
