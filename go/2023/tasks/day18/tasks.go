//go:generate go run aoc/cmd/generate-tests

package day18

import (
	"aoc/framework/geometry/geo2d"
	_math "aoc/framework/math"
	"aoc/framework/tasks"
	"go/types"
	"math"
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	instructions := parse1(data)
	vertices := findVertices(instructions, geo2d.Point{})
	result.Value = calculateArea(vertices)

	return
}

func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	instructions := parse2(data)
	vertices := findVertices(instructions, geo2d.Point{})
	result.Value = calculateArea(vertices)

	return
}

func findVertices(instructions []Instruction, start geo2d.Point) []geo2d.Point {
	numVertices := len(instructions) + 1

	corners := make([]geo2d.Point, numVertices)
	corners[0] = start

	for index, instruction := range instructions {
		current := corners[index]

		switch instruction.direction {
		case geo2d.North:
			corners[index+1] = geo2d.Point{X: current.X, Y: current.Y - instruction.meters}
		case geo2d.East:
			corners[index+1] = geo2d.Point{X: current.X + instruction.meters, Y: current.Y}
		case geo2d.South:
			corners[index+1] = geo2d.Point{X: current.X, Y: current.Y + instruction.meters}
		case geo2d.West:
			corners[index+1] = geo2d.Point{X: current.X - instruction.meters, Y: current.Y}
		default:
			panic("invalid direction")
		}
	}

	vertices := make([]geo2d.Point, numVertices)

	for currentIndex := 0; currentIndex < numVertices; currentIndex++ {
		previousIndex := currentIndex - 1
		if previousIndex < 0 {
			previousIndex = numVertices - 1
		}

		nextIndex := currentIndex + 1
		if nextIndex >= numVertices {
			nextIndex = 0
		}

		previous := corners[previousIndex]
		current := corners[currentIndex]
		next := corners[nextIndex]

		incomingDirection, _ := previous.OrientationOf(current)
		outgoingDirection, _ := current.OrientationOf(next)

		switch {
		case incomingDirection == geo2d.North && outgoingDirection == geo2d.East:
			vertices[currentIndex] = geo2d.Point{X: current.X, Y: current.Y}
		case incomingDirection == geo2d.North && outgoingDirection == geo2d.West:
			vertices[currentIndex] = geo2d.Point{X: current.X, Y: current.Y + 1}
		case incomingDirection == geo2d.East && outgoingDirection == geo2d.North:
			vertices[currentIndex] = geo2d.Point{X: current.X, Y: current.Y}
		case incomingDirection == geo2d.East && outgoingDirection == geo2d.South:
			vertices[currentIndex] = geo2d.Point{X: current.X + 1, Y: current.Y}
		case incomingDirection == geo2d.South && outgoingDirection == geo2d.East:
			vertices[currentIndex] = geo2d.Point{X: current.X + 1, Y: current.Y}
		case incomingDirection == geo2d.South && outgoingDirection == geo2d.West:
			vertices[currentIndex] = geo2d.Point{X: current.X + 1, Y: current.Y + 1}
		case incomingDirection == geo2d.West && outgoingDirection == geo2d.North:
			vertices[currentIndex] = geo2d.Point{X: current.X, Y: current.Y + 1}
		case incomingDirection == geo2d.West && outgoingDirection == geo2d.South:
			vertices[currentIndex] = geo2d.Point{X: current.X, Y: current.Y + 1}
		}
	}

	return vertices
}

func calculateArea(vertices []geo2d.Point) int {
	var totalArea int

	maxY := lo.MaxBy(vertices, func(a geo2d.Point, b geo2d.Point) bool {
		return a.Y > b.Y
	}).Y

	for currentIndex := 0; currentIndex < len(vertices); currentIndex++ {
		nextIndex := currentIndex + 1
		if nextIndex >= len(vertices) {
			nextIndex = 0
		}

		current := vertices[currentIndex]
		next := vertices[nextIndex]

		if current.X == next.X {
			continue
		}

		width := _math.AbsInt(current.X - next.X)
		area := width * (maxY - current.Y)

		switch {
		case current.X < next.X:
			totalArea += area
		case current.X > next.X:
			totalArea -= area
		}
	}

	return int(math.Abs(float64(totalArea)))
}

type Instruction struct {
	direction geo2d.Orientation
	meters    int
}

var instructionPattern1 = regexp.MustCompile(`(?m)^([LRUD]) (\d+) .*$`)
var instructionPattern2 = regexp.MustCompile(`(?m)^[LRUD] \d+ \(#(.{5})(\d)\)$`)

func parse1(data string) (instructions []Instruction) {
	matches := instructionPattern1.FindAllStringSubmatch(data, -1)

	for _, match := range matches {
		instruction := Instruction{}

		switch match[1] {
		case "U":
			instruction.direction = geo2d.North
		case "R":
			instruction.direction = geo2d.East
		case "D":
			instruction.direction = geo2d.South
		case "L":
			instruction.direction = geo2d.West
		default:
			panic("invalid direction")
		}

		if meters, err := strconv.Atoi(match[2]); err != nil {
			panic(err)
		} else {
			instruction.meters = meters
		}

		instructions = append(instructions, instruction)
	}
	return
}

func parse2(data string) (instructions []Instruction) {
	matches := instructionPattern2.FindAllStringSubmatch(data, -1)

	for _, match := range matches {
		instruction := Instruction{}

		if meters, err := strconv.ParseUint(match[1], 16, 64); err != nil {
			panic(err)
		} else {
			instruction.meters = int(meters)
		}

		switch match[2] {
		case "0":
			instruction.direction = geo2d.East
		case "1":
			instruction.direction = geo2d.South
		case "2":
			instruction.direction = geo2d.West
		case "3":
			instruction.direction = geo2d.North
		default:
			panic("invalid direction")
		}

		instructions = append(instructions, instruction)
	}
	return
}
