package day10

import (
	"2023/src/tasks/day10/model"
	"aoc/framework/tasks"
	"go/types"

	lop "github.com/samber/lo/parallel"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	surface := model.NewSurface(data)

	loop := surface.FindLoop()

	result.Value = len(loop) / 2

	return
}

func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	surface := model.NewSurface(data)

	loop := surface.FindLoop()
	lop.ForEach(loop, func(item model.Segment, index int) {
		item.Pipe.PartOfLoop = true
	})

	result.Value = len(surface.FindAllPointsInsideLoop(loop))

	return
}
