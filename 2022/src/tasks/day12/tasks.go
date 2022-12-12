package day12

import (
	"2022/src/framework"
	"2022/src/framework/test"
	"2022/src/tasks/day12/model"
	"log"
	"sort"
)

func Task1(filePath string) (result test.TaskResult[int]) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	grid := model.ReadGrid(data)

	cameFrom, err := grid.PathTo(grid.Start, grid.Destination)
	if err != nil {
		log.Fatal(err)
	}

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
	data, err := framework.ReadLines(filePath)
	if err != nil {
		result.Error = err
		return
	}

	grid := model.ReadGrid(data)

	var startingPoints []model.Position
	for x := range grid.Squares {
		for y := range grid.Squares[x] {
			if grid.Squares[x][y] == 1 {
				startingPoints = append(startingPoints, model.Position{X: x, Y: y})
			}
		}
	}

	var pathLengths []int

	for index, startingPoint := range startingPoints {
		log.Printf("calculating length for starting point %v of %v\n", index+1, len(startingPoints))
		cameFrom, err := grid.PathTo(startingPoint, grid.Destination)
		if err != nil {
			log.Println(err)
			continue
		}

		current := grid.Destination
		path := make([]model.Position, 0)
		for current != startingPoint {
			path = append(path, current)
			current = cameFrom[current]
		}
		path = append(path, startingPoint)

		pathLengths = append(pathLengths, len(path)-1)
	}

	sort.Sort(sort.IntSlice(pathLengths))

	result.Value = pathLengths[0]

	return
}
