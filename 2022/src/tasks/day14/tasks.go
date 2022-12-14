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

	spilling := false
	for !spilling {
		if cave.DropSand(framework.Point{X: 500, Y: 0}) {
			result.Value++
		} else {
			spilling = true
		}
	}

	framework.DrawPointGrid(cave.Area, map[model.Material]rune{
		model.ROCK: '#',
		model.SAND: 'o',
		model.AIR:  '.',
	})

	return
}

func Task2(filePath string) (result test.TaskResult[int]) {
	_, err := framework.ReadLineBlocks(filePath)
	if err != nil {
		result.Error = err
		return
	}

	return
}
