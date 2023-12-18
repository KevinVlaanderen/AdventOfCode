package day11

import (
	"2023/src/framework"
	"2023/src/framework/geometry"
	"2023/src/framework/geometry/grid"
	"2023/src/framework/math"
	"github.com/samber/lo"
)

type Galaxy struct {
	id int
}

func Task1(data string) (result framework.Result[int]) {
	universe := NewUniverse(framework.Lines(data))

	result.Value = universe.CalculateDistances(2)

	return
}

func Task2(data string) (result framework.Result[int]) {
	universe := NewUniverse(framework.Lines(data))

	result.Value = universe.CalculateDistances(1000000)

	return
}

type Universe grid.SparseGrid[bool]

func NewUniverse(data []string) Universe {
	universe := grid.NewSparseGrid[bool]()

	for y, line := range data {
		for x, char := range line {
			if char == '#' {
				universe.Add(geometry.Point{X: x, Y: y}, true)
			}
		}
	}
	return Universe(universe)
}

func (universe Universe) CalculateDistances(factor int) int {
	var result int
	g := grid.SparseGrid[bool](universe)

	xMin, xMax, yMin, yMax := g.Boundaries()
	var emptyRows, emptyCols []int

	for x := xMin + 1; x < xMax; x++ {
		colEmpty := true
		for y := yMin; y <= yMax; y++ {
			if _, found := g.Get(geometry.Point{X: x, Y: y}); found {
				colEmpty = false
				break
			}
		}
		if colEmpty {
			emptyCols = append(emptyCols, x)
		}
	}
	for y := yMin + 1; y < yMax; y++ {
		rowEmpty := true
		for x := xMin; x <= xMax; x++ {
			if _, found := g.Get(geometry.Point{X: x, Y: y}); found {
				rowEmpty = false
				break
			}
		}
		if rowEmpty {
			emptyRows = append(emptyRows, y)
		}
	}

	galaxies := g.Keys()
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
