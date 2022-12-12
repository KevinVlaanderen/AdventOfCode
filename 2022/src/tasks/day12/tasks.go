package day12

import (
	"2022/src/framework"
	"2022/src/framework/test"
	"2022/src/tasks/day12/model"
)

func Task1(filePath string) (result test.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	grid := model.ReadGrid(data)

	cameFrom := grid.PathTo(grid.Destination)

	current := grid.Destination
	path := make([]model.Position, 0)

	for current != grid.Start {
		path = append(path, cameFrom[current])
		current = cameFrom[current]
	}
	path = append(path, grid.Start)

	result.Value = len(path) - 1

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
