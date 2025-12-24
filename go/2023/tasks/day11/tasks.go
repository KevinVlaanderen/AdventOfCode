//go:generate go run aoc/cmd/generate-tests

package day11

import (
	"aoc/framework/geometry/geo2d"
	"aoc/framework/geometry/geo2d/grid"
	"aoc/framework/math"
	"aoc/framework/tasks"
	"go/types"

	"github.com/samber/lo"
)

type Galaxy struct {
	id int
}

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	universe := NewUniverse(tasks.Lines(data))

	result.Value = universe.CalculateDistances(2)

	return
}

func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	universe := NewUniverse(tasks.Lines(data))

	result.Value = universe.CalculateDistances(1000000)

	return
}

type Universe struct {
	grid grid.Grid[bool]
}

func NewUniverse(data []string) Universe {
	universe := grid.NewSparseGrid[bool]()

	for y, line := range data {
		for x, char := range line {
			if char == '#' {
				universe.Set(geo2d.Point{X: x, Y: y}, true)
			}
		}
	}
	return Universe{universe}
}

func (universe Universe) CalculateDistances(factor int) int {
	var result int

	minX, minY, maxX, maxY := universe.grid.Bounds()
	var emptyRows, emptyCols []int

	for x := minX + 1; x < maxX; x++ {
		colEmpty := true
		for y := minY; y <= maxY; y++ {
			if _, found := universe.grid.Get(geo2d.Point{X: x, Y: y}); found {
				colEmpty = false
				break
			}
		}
		if colEmpty {
			emptyCols = append(emptyCols, x)
		}
	}
	for y := minY + 1; y < maxY; y++ {
		rowEmpty := true
		for x := minX; x <= maxX; x++ {
			if _, found := universe.grid.Get(geo2d.Point{X: x, Y: y}); found {
				rowEmpty = false
				break
			}
		}
		if rowEmpty {
			emptyRows = append(emptyRows, y)
		}
	}

	galaxies := lo.ChannelToSlice(universe.grid.PointsSet())

	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			currentPoint, nextPoint := galaxies[i], galaxies[j]

			emptyColsBetween := lo.CountBy(emptyCols, func(item int) bool {
				return item > currentPoint.X && item < nextPoint.X || item > nextPoint.X && item < currentPoint.X
			})
			emptyRowsBetween := lo.CountBy(emptyRows, func(item int) bool {
				return item > currentPoint.Y && item < nextPoint.Y || item > nextPoint.Y && item < currentPoint.Y
			})

			distance := math.AbsInt(nextPoint.X-currentPoint.X) + math.AbsInt(nextPoint.Y-currentPoint.Y)
			value := distance + (emptyColsBetween * (factor - 1)) + (emptyRowsBetween * (factor - 1))

			result += value
		}
	}
	return result
}
