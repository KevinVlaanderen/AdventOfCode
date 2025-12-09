package day6

import (
	"aoc/2022/tasks/day6/model"
	"aoc/framework/tasks"
	"go/types"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	signal := parse(data)

	result.Value = signal.FindMarker(4)

	return
}

func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	signal := parse(data)

	result.Value = signal.FindMarker(14)

	return
}

func parse(data string) model.Signal {
	return model.Signal(data)
}
