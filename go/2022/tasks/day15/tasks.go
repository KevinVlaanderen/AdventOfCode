//go:generate go run aoc/cmd/generate-tests

package day15

import (
	"aoc/2022/tasks/day15/model"
	"aoc/framework"
	"aoc/framework/geometry/geo2d"
	"aoc/framework/math"
	"aoc/framework/tasks"
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

// Task1 type:mock 	file:data	param:10		expected:26
// Task1 type:real 	file:day15 	param:2000000	expected:4886370
func Task1(data string, testRow int) (result tasks.Result[int]) {
	sensors := parse(data)

	var initialized bool
	var xMin, xMax int

	for _, sensor := range sensors {
		diffX := math.AbsInt(sensor.X - sensor.ClosestBeacon.X)
		diffY := math.AbsInt(sensor.Y - sensor.ClosestBeacon.Y)
		distance := diffX + diffY

		if !initialized || sensor.X-distance < xMin {
			xMin = sensor.X - distance
		}
		if !initialized || sensor.X+distance > xMax {
			xMax = sensor.X + distance
		}
		initialized = true
	}

	for _, x := range framework.Range(xMin, xMax-xMin+1, 1) {
		testPoint := geo2d.Point{X: x, Y: testRow}

		canContainBeacon := true
		for _, sensor := range sensors {
			diffX := math.AbsInt(sensor.X - testPoint.X)
			diffY := math.AbsInt(sensor.Y - testPoint.Y)
			distance := diffX + diffY

			if testPoint == sensor.ClosestBeacon {
				break
			} else if distance <= sensor.Distance {
				canContainBeacon = false
				break
			}
		}
		if !canContainBeacon {
			result.Value++
		}
	}

	return
}

var beaconPattern = regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

func parse(data string) []model.Sensor {
	return lo.Map(tasks.Lines(data), func(line string, index int) model.Sensor {
		matches := beaconPattern.FindStringSubmatch(line)
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		beaconX, _ := strconv.Atoi(matches[3])
		beaconY, _ := strconv.Atoi(matches[4])

		return model.NewSensor(geo2d.Point{X: x, Y: y}, geo2d.Point{X: beaconX, Y: beaconY})
	})
}
