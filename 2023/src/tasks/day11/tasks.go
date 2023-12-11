package day11

import (
	"2023/src/framework"
	"2023/src/framework/geometry"
	"2023/src/framework/math"
	"github.com/samber/lo"
)

type Galaxy struct {
	id int
}

func Task1(filePath string) (result framework.Result[int]) {
	universe := geometry.NewGrid[Galaxy]()

	var x, y int
	n := 1
	for line := range framework.ReadLines(filePath) {
		for _, char := range line {
			if char == '#' {
				universe.Add(geometry.Point{X: x, Y: y}, Galaxy{n})
				n++
			}
			x++
		}
		y++
		x = 0
	}

	xMin, xMax, yMin, yMax := universe.Boundaries()
	var emptyRows, emptyCols []int

	for x = xMin + 1; x < xMax; x++ {
		colEmpty := true
		for y = yMin; y <= yMax; y++ {
			if _, found := universe.Get(geometry.Point{X: x, Y: y}); found {
				colEmpty = false
				break
			}
		}
		if colEmpty {
			emptyCols = append(emptyCols, x)
		}
	}
	for y = yMin + 1; y < yMax; y++ {
		rowEmpty := true
		for x = xMin; x <= xMax; x++ {
			if _, found := universe.Get(geometry.Point{X: x, Y: y}); found {
				rowEmpty = false
				break
			}
		}
		if rowEmpty {
			emptyRows = append(emptyRows, y)
		}
	}

	galaxies := universe.Keys()
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
			value := distance + emptyColsBetween + emptyRowsBetween
			result.Value += value
		}
	}

	return
}
