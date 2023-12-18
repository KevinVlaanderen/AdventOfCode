package day18

import (
	"2023/src/framework"
	"2023/src/framework/geometry"
	"2023/src/framework/geometry/grid"
	"github.com/oleiade/lane/v2"
	"regexp"
	"strconv"
)

func Task1(data string) (result framework.Result[int]) {
	instructions := parse(data)

	lagoon := grid.NewSparseGrid[bool]()

	current := geometry.Point{}
	for _, instruction := range instructions {
		for i := 0; i < instruction.meters; i++ {
			current = current.Neighbour(instruction.direction)
			lagoon.Add(current, true)
		}
	}
	fill(&lagoon, geometry.Point{X: current.X + 1, Y: current.Y + 1})

	lagoon.DrawPointGrid(map[bool]rune{true: '#'}, '.')

	result.Value = len(lagoon.Keys())

	return
}

func fill(g *grid.SparseGrid[bool], start geometry.Point) {
	stack := lane.NewStack(start)

	for stack.Size() > 0 {
		current, _ := stack.Pop()
		g.Add(current, true)

		for _, offset := range current.NeighbourOffsets(geometry.Orthogonal) {
			next := geometry.Point{X: current.X + offset.X, Y: current.Y + offset.Y}
			if !g.InBounds(next) {
				continue
			}
			if _, found := g.Get(next); found {
				continue
			}
			stack.Push(next)
		}
	}
}

type Instruction struct {
	direction geometry.Orientation
	meters    int
}

var instructionPattern = regexp.MustCompile(`(?m)^([LRUD]) (\d+) .*$`)

func parse(data string) (instructions []Instruction) {
	matches := instructionPattern.FindAllStringSubmatch(data, -1)

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
