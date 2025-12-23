//go:generate go run aoc/cmd/generate-tests

package day10

import (
	"aoc/2023/tasks/day10/model"
	"aoc/framework/tasks"
	"go/types"

	lop "github.com/samber/lo/parallel"
)

// Task1 type:mock 	file:data	expected:4
// Task1 type:mock 	file:data2	expected:8
// Task1 type:real 	file:day10 	expected:6757
func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	surface := model.NewSurface(data)

	loop := surface.FindLoop()

	result.Value = len(loop) / 2

	return
}

// Task2 type:mock 	file:data3	expected:4
// Task2 type:mock 	file:data4	expected:8
// Task2 type:mock 	file:data5	expected:10
// Task2 type:real 	file:day10 	expected:523
func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	surface := model.NewSurface(data)

	loop := surface.FindLoop()
	lop.ForEach(loop, func(item model.Segment, index int) {
		item.Pipe.PartOfLoop = true
	})

	result.Value = len(surface.FindAllPointsInsideLoop(loop))

	return
}
