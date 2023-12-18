package day18

import (
	"2023/src/framework"
	"2023/src/framework/geometry"
	_math "2023/src/framework/math"
	"github.com/samber/lo"
	"math"
	"regexp"
	"strconv"
)

func Task1(data string) (result framework.Result[int64]) {
	instructions := parse1(data)
	vertices := findVertices(instructions, geometry.Point{})
	result.Value = calculateArea(vertices)

	return
}

func Task2(data string) (result framework.Result[int64]) {
	instructions := parse2(data)
	vertices := findVertices(instructions, geometry.Point{})
	result.Value = calculateArea(vertices)

	return
}

func findVertices(instructions []Instruction, start geometry.Point) []geometry.Point {
	numVertices := len(instructions) + 1

	corners := make([]geometry.Point, numVertices)
	corners[0] = start

	for index, instruction := range instructions {
		current := corners[index]

		switch instruction.direction {
		case geometry.North:
			corners[index+1] = geometry.Point{X: current.X, Y: current.Y - instruction.meters}
		case geometry.East:
			corners[index+1] = geometry.Point{X: current.X + instruction.meters, Y: current.Y}
		case geometry.South:
			corners[index+1] = geometry.Point{X: current.X, Y: current.Y + instruction.meters}
		case geometry.West:
			corners[index+1] = geometry.Point{X: current.X - instruction.meters, Y: current.Y}
		default:
			panic("invalid direction")
		}
	}

	vertices := make([]geometry.Point, numVertices)

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
		case incomingDirection == geometry.North && outgoingDirection == geometry.East:
			vertices[currentIndex] = geometry.Point{X: current.X, Y: current.Y}
		case incomingDirection == geometry.North && outgoingDirection == geometry.West:
			vertices[currentIndex] = geometry.Point{X: current.X, Y: current.Y + 1}
		case incomingDirection == geometry.East && outgoingDirection == geometry.North:
			vertices[currentIndex] = geometry.Point{X: current.X, Y: current.Y}
		case incomingDirection == geometry.East && outgoingDirection == geometry.South:
			vertices[currentIndex] = geometry.Point{X: current.X + 1, Y: current.Y}
		case incomingDirection == geometry.South && outgoingDirection == geometry.East:
			vertices[currentIndex] = geometry.Point{X: current.X + 1, Y: current.Y}
		case incomingDirection == geometry.South && outgoingDirection == geometry.West:
			vertices[currentIndex] = geometry.Point{X: current.X + 1, Y: current.Y + 1}
		case incomingDirection == geometry.West && outgoingDirection == geometry.North:
			vertices[currentIndex] = geometry.Point{X: current.X, Y: current.Y + 1}
		case incomingDirection == geometry.West && outgoingDirection == geometry.South:
			vertices[currentIndex] = geometry.Point{X: current.X, Y: current.Y + 1}
		}
	}

	return vertices
}

func calculateArea(vertices []geometry.Point) int64 {
	var totalArea int64

	maxY := lo.MaxBy(vertices, func(a geometry.Point, b geometry.Point) bool {
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
		area := int64(width * (maxY - current.Y))

		switch {
		case current.X < next.X:
			totalArea += area
		case current.X > next.X:
			totalArea -= area
		}
	}

	return int64(math.Abs(float64(totalArea)))
}

type Instruction struct {
	direction geometry.Orientation
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
			instruction.direction = geometry.North
		case "R":
			instruction.direction = geometry.East
		case "D":
			instruction.direction = geometry.South
		case "L":
			instruction.direction = geometry.West
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
			instruction.direction = geometry.East
		case "1":
			instruction.direction = geometry.South
		case "2":
			instruction.direction = geometry.West
		case "3":
			instruction.direction = geometry.North
		default:
			panic("invalid direction")
		}

		instructions = append(instructions, instruction)
	}
	return
}
