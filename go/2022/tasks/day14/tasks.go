package day14

import (
	"aoc/2022/tasks/day14/model"
	"aoc/framework/geometry/geo2d"
	"aoc/framework/geometry/geo2d/grid"
	"aoc/framework/tasks"
	"go/types"
	"strconv"
	"strings"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	cave := parse(data)

	for cave.DropSand(geo2d.Point{X: 500, Y: 0}) {
		result.Value++
	}

	return
}

func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	cave := parse(data)

	cave.AddRock(
		geo2d.Point{X: cave.MinX - cave.MaxY, Y: cave.MaxY + 2},
		geo2d.Point{X: cave.MaxX + cave.MaxY, Y: cave.MaxY + 2})

	for cave.DropSand(geo2d.Point{X: 500, Y: 0}) {
		result.Value++
	}

	return
}

func parse(data string) *model.Cave {
	cave := model.Cave{Area: grid.NewSparseGrid[model.Material]()}

	for _, line := range tasks.Lines(data) {
		lineParts := strings.Split(line, " -> ")

		var positions []geo2d.Point
		for _, linePart := range lineParts {
			rockParts := strings.Split(linePart, ",")
			x, _ := strconv.Atoi(rockParts[0])
			y, _ := strconv.Atoi(rockParts[1])
			positions = append(positions, geo2d.Point{X: x, Y: y})
		}

		for index := range positions {
			if index == len(positions)-1 {
				break
			}

			cave.AddRock(positions[index], positions[index+1])
		}
	}

	return &cave
}
