package day14

import (
	"2022/src/framework/geometry"
	"2022/src/framework/tasks"
	"2022/src/tasks/day14/model"
	"strconv"
	"strings"
)

func Task1(filePath string) (result tasks.TaskResult[int]) {
	cave := <-tasks.ReadStream(filePath, createParser())

	for cave.DropSand(geometry.Point{X: 500, Y: 0}) {
		result.Value++
	}

	geometry.DrawPointGrid(cave.Area, map[model.Material]rune{
		model.ROCK: '#',
		model.SAND: 'o',
	}, '.')

	return
}

func Task2(filePath string) (result tasks.TaskResult[int]) {
	cave := <-tasks.ReadStream(filePath, createParser())

	cave.AddRock(
		geometry.Point{X: cave.MinX - cave.MaxY, Y: cave.MaxY + 2},
		geometry.Point{X: cave.MaxX + cave.MaxY, Y: cave.MaxY + 2})

	for cave.DropSand(geometry.Point{X: 500, Y: 0}) {
		result.Value++
	}

	geometry.DrawPointGrid(cave.Area, map[model.Material]rune{
		model.ROCK: '#',
		model.SAND: 'o',
	}, '.')

	return
}

func createParser() func(line string) (result *model.Cave, hasResult bool, err error) {
	cave := model.Cave{Area: map[geometry.Point]model.Material{}}

	return func(line string) (result *model.Cave, hasResult bool, err error) {
		if line == "" {
			return &cave, true, err
		}

		lineParts := strings.Split(line, " -> ")

		var positions []geometry.Point
		for _, linePart := range lineParts {
			rockParts := strings.Split(linePart, ",")
			x, _ := strconv.Atoi(rockParts[0])
			y, _ := strconv.Atoi(rockParts[1])
			positions = append(positions, geometry.Point{X: x, Y: y})
		}

		for index := range positions {
			if index == len(positions)-1 {
				break
			}

			cave.AddRock(positions[index], positions[index+1])
		}

		return result, false, nil
	}
}