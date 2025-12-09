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

type Rectangle struct {
	U, V geo2d.Point
}

type Area struct {
	Rectangle Rectangle
	Size      int
}

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	points := parse(data)

	areas := calculateAreas(points)

	result.Value = areas[0].Size

	return
}

func calculateAreas(points []geo2d.Point) (areas []Area) {
	areaMap := make(map[Rectangle]int)

	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			size := (1 + math.AbsInt(points[i].X-points[j].X)) * (1 + math.AbsInt(points[i].Y-points[j].Y))
			areaMap[Rectangle{U: points[i], V: points[j]}] = size
		}
	}

	for k, v := range areaMap {
		areas = append(areas, Area{k, v})
	}

	sort.Slice(areas, func(i, j int) bool {
		return areas[i].Size > areas[j].Size
	})

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
