package day15

import (
	"2022/src/framework/generators"
	"2022/src/framework/geometry"
	"2022/src/framework/math"
	"2022/src/framework/tasks"
	"2022/src/tasks/day15/model"
	"regexp"
	"strconv"
)

func CreateTask1(testRow int) func(filePath string) (result tasks.TaskResult[int]) {
	return func(filePath string) (result tasks.TaskResult[int]) {
		sensors := tasks.Read(filePath, parser)

		var initialized bool
		var xMin, xMax int

		for _, sensor := range sensors {
			diffX := math.AbsInt(sensor.X - sensor.ClosestBeacon.X)
			diffY := math.AbsInt(sensor.Y - sensor.ClosestBeacon.Y)
			distance := diffX + diffY

			if !initialized || sensor.X - distance < xMin {
				xMin = sensor.X - distance
			}
			if !initialized || sensor.X + distance > xMax {
				xMax = sensor.X + distance
			}
			initialized = true
		}

		for _, x := range generators.Range(xMin, xMax-xMin+1, 1) {
			testPoint := geometry.Point{X: x, Y: testRow}

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
}

var beaconPattern = regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

func parser(line string) (result model.Sensor, hasResult bool, err error) {
	if line == "" {
		return result, false, nil
	}

	matches := beaconPattern.FindStringSubmatch(line)
	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])
	beaconX, _ := strconv.Atoi(matches[3])
	beaconY, _ := strconv.Atoi(matches[4])

	return model.NewSensor(geometry.Point{X: x, Y: y}, geometry.Point{X: beaconX, Y: beaconY}), true, nil
}
