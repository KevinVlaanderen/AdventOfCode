package day14

import (
	"2022/src/framework"
	"2022/src/framework/test"
	"2022/src/tasks/day14/model"
)

func Task1(filePath string) (result test.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	cave := model.NewCave(data)

	for cave.DropSand(framework.Point{X: 500, Y: 0}) {
		result.Value++
	}

	framework.DrawPointGrid(cave.Area, map[model.Material]rune{
		model.ROCK: '#',
		model.SAND: 'o',
	}, '.')

	return
}

func Task2(filePath string) (result test.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	cave := model.NewCave(data)
	cave.AddRock(
		framework.Point{X: cave.MinX - cave.MaxY, Y: cave.MaxY + 2},
		framework.Point{X: cave.MaxX + cave.MaxY, Y: cave.MaxY + 2})

	for cave.DropSand(framework.Point{X: 500, Y: 0}) {
		result.Value++
	}

	framework.DrawPointGrid(cave.Area, map[model.Material]rune{
		model.ROCK: '#',
		model.SAND: 'o',
	}, '.')

	return
}
