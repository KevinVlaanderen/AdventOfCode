package day10

import (
	"2023/src/framework"
	"2023/src/tasks/day10/model"
)

func Task1(filePath string) (result framework.Result[int]) {
	data := framework.ReadAll(filePath)
	surface := model.NewSurface(data)

	loop := surface.FindLoop()

	result.Value = len(loop) / 2

	return
}

func Task2(filePath string) (result framework.Result[int]) {
	data := framework.ReadAll(filePath)
	surface := model.NewSurface(data)

	loop := surface.FindLoop()
	found := surface.FindAllPointsInsideLoop(loop)

	result.Value = len(found)

	return
}
