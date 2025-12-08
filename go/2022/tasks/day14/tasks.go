package day14

import (
	"aoc/2022/tasks/day14/model"
	"aoc/framework"
	"aoc/framework/geometry"
	"aoc/framework/geometry/grid"
	"go/types"
	"strconv"
	"strings"
)

func Task1(data string, _ types.Nil) (result framework.Result[int]) {
	cave := parse(data)

	for cave.DropSand(geometry.Point{X: 500, Y: 0}) {
		result.Value++
	}

	// cave.Area.DrawPointGrid(func(material *model.Material, x int, y int) (rune, bool) {
	// 	switch *material {
	// 	case model.ROCK:
	// 		return '#', true
	// 	case model.SAND:
	// 		return 'o', true
	// 	default:
	// 		return '.', false
	// 	}
	// }, map[model.Material]rune{})

	return
}

func Task2(data string, _ types.Nil) (result framework.Result[int]) {
	cave := parse(data)

	cave.AddRock(
		geometry.Point{X: cave.MinX - cave.MaxY, Y: cave.MaxY + 2},
		geometry.Point{X: cave.MaxX + cave.MaxY, Y: cave.MaxY + 2})

	for cave.DropSand(geometry.Point{X: 500, Y: 0}) {
		result.Value++
	}

	// cave.Area.DrawPointGrid(func(material *model.Material, x int, y int) (rune, bool) {
	// 	switch *material {
	// 	case model.ROCK:
	// 		return '#', true
	// 	case model.SAND:
	// 		return 'o', true
	// 	default:
	// 		return '.', false
	// 	}
	// }, map[model.Material]rune{})

	return
}

func parse(data string) *model.Cave {
	lines := framework.Lines(data)
	width := len(lines[0])
	height := len(lines)

	cave := model.Cave{Area: grid.NewGrid[model.Material](width, height)}

	for _, line := range framework.Lines(data) {
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
	}

	return &cave
}
