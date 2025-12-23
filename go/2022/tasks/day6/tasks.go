//go:generate go run aoc/cmd/generate-tests

package day6

import (
	"aoc/2022/tasks/day6/model"
	"aoc/framework/tasks"
	"go/types"
)

// Task1 type:mock 	file:data1	expected:7
// Task1 type:mock 	file:data2	expected:5
// Task1 type:mock 	file:data3	expected:6
// Task1 type:mock 	file:data4	expected:10
// Task1 type:mock 	file:data5	expected:11
// Task1 type:real 	file:day6 	expected:1042
func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	signal := parse(data)

	result.Value = signal.FindMarker(4)

	return
}

// Task2 type:mock 	file:data1	expected:19
// Task2 type:mock 	file:data2	expected:23
// Task2 type:mock 	file:data3	expected:23
// Task2 type:mock 	file:data4	expected:29
// Task2 type:mock 	file:data5	expected:26
// Task2 type:real 	file:day6 	expected:2980
func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	signal := parse(data)

	result.Value = signal.FindMarker(14)

	return
}

func parse(data string) model.Signal {
	return model.Signal(data)
}
