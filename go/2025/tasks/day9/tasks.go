//go:generate go run aoc/cmd/generate-tests

package day9

import (
	"aoc/framework/geometry/geo2d"
	"aoc/framework/math"
	"aoc/framework/tasks"
	"go/types"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

// Task1 type:mock 	file:data	expected:50
// Task1 type:real 	file:day9 	expected:4777409595
func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	points := parse(data)

	areas := calculateAreas(points)
	sort.Slice(areas, func(i, j int) bool {
		return areas[i].Size() > areas[j].Size()
	})

	result.Value = areas[0].Size()

	return
}

// Task2 type:mock 	file:data	expected:24
// Task2 type:real 	file:day9 	expected:1473551379
func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	points := parse(data)

	lines := createLines(points)
	areas := calculateAreas(points)
	sort.Slice(areas, func(i, j int) bool {
		return areas[i].Size() > areas[j].Size()
	})

	type State struct {
		line              geo2d.Line
		previousLineIndex int
		previousCoord     int
	}

	var largestArea geo2d.Area

largestAreaLoop:
	for _, largestArea = range areas {
		minX := math.MinInt(largestArea.U.X, largestArea.V.X)
		maxX := math.MaxInt(largestArea.U.X, largestArea.V.X)
		minY := math.MinInt(largestArea.U.Y, largestArea.V.Y)
		maxY := math.MaxInt(largestArea.U.Y, largestArea.V.Y)

		for _, line := range lines {
			lineMinX := math.MinInt(line.A.X, line.B.X)
			lineMaxX := math.MaxInt(line.A.X, line.B.X)
			lineMinY := math.MinInt(line.A.Y, line.B.Y)
			lineMaxY := math.MaxInt(line.A.Y, line.B.Y)

			if minX < lineMaxX && maxX > lineMinX && minY < lineMaxY && maxY > lineMinY {
				continue largestAreaLoop
			}
		}

		break largestAreaLoop
	}

	result.Value = largestArea.Size()

	return
}

func parse(data string) []geo2d.Point {
	return lo.Map(tasks.Lines(data), func(line string, index int) geo2d.Point {
		parts := strings.Split(line, ",")
		partsInt := lo.Map(parts, func(part string, index int) int {
			n, _ := strconv.Atoi(part)
			return n
		})
		return geo2d.Point{X: partsInt[0], Y: partsInt[1]}
	})
}

func createLines(points []geo2d.Point) (lines []geo2d.Line) {
	for i := 0; i < len(points)-1; i++ {
		lines = append(lines, geo2d.Line{A: points[i], B: points[i+1]})
	}
	lines = append(lines, geo2d.Line{A: points[len(points)-1], B: points[0]})
	return
}

func calculateAreas(points []geo2d.Point) (areas []geo2d.Area) {
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			areas = append(areas, geo2d.NewArea(points[i], points[j]))
		}
	}
	return
}
