//go:generate go run aoc/cmd/generate-tests

package day12

import (
	"aoc/2022/tasks/day12/model"
	"aoc/framework/tasks"
	"go/types"
	"sort"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	grid := parse(data)

	cameFrom, err := grid.PathTo(grid.Start, grid.Destination)
	if err != nil {
		panic(err)
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

func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	grid := parse(data)

	var startingPoints []model.Position
	for x := range grid.Squares {
		for y := range grid.Squares[x] {
			if grid.Squares[x][y] == 1 {
				startingPoints = append(startingPoints, model.Position{X: x, Y: y})
			}
		}
	}

	var pathLengths []int

	for _, startingPoint := range startingPoints {
		cameFrom, err := grid.PathTo(startingPoint, grid.Destination)
		if err != nil {
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

func parse(data string) *model.Grid {
	lines := tasks.Lines(data)

	grid := model.Grid{}
	grid.Squares = make([][]int, len(lines[0]))
	for i := range grid.Squares {
		grid.Squares[i] = make([]int, len(lines))
	}

	for y := range lines {
		for x, char := range []rune(lines[y]) {
			switch char {
			case 'S':
				grid.Start = model.Position{X: x, Y: y}
				grid.Squares[x][y] = 1
			case 'E':
				grid.Destination = model.Position{X: x, Y: y}
				grid.Squares[x][y] = 26
			default:
				grid.Squares[x][y] = int(char) - 96
			}
		}
	}

	return &grid
}
