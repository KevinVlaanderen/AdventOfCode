package day10

import (
	"2023/src/framework"
	"2023/src/framework/geometry"
	"2023/src/tasks/day10/model"
	"github.com/samber/lo"
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
	loopPoints := lo.Map(loop, func(item model.Segment, index int) geometry.Point {
		return *item.Point
	})
	for _, point := range surface.Grid.Keys() {
		if !lo.Contains(loopPoints, point) {
			surface.Grid.Delete(point)
		}
	}
	found := surface.FindAllPointsInsideLoop(loop)

	result.Value = len(found)

	return
}
