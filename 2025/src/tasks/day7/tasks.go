package day7

import (
	"2025/src/framework"
	"2025/src/framework/geometry"
	"2025/src/framework/geometry/grid"
	"go/types"

	"github.com/oleiade/lane/v2"
)

type SpaceType uint8

const (
	Start SpaceType = iota
	Empty
	Splitter
	Beam
)

func Task1(data string, _ types.Nil) (result framework.Result[int]) {
	lab, start := parse(data)
	beams := lane.NewQueue[geometry.Point](start.Neighbour(geometry.South))

	for beams.Size() > 0 {
		currentBeam, _ := beams.Head()
		nextPosition := currentBeam.Neighbour(geometry.South)
		spaceType, found := lab.Get(&nextPosition)
		if found {
			if spaceType == Empty {
				_ = lab.Set(&nextPosition, Beam)
				beams.Enqueue(nextPosition)
			} else if spaceType == Splitter {
				result.Value++

				left := nextPosition.Neighbour(geometry.West)
				_ = lab.Set(&left, Beam)
				beams.Enqueue(left)

				right := nextPosition.Neighbour(geometry.East)
				_ = lab.Set(&right, Beam)
				beams.Enqueue(right)
			}
		}
		beams.Dequeue()
	}

	return
}

func parse(data string) (space grid.Grid[SpaceType], start geometry.Point) {
	lines := framework.CharLines(data)
	width := len(lines[0])
	height := len(lines)

	space = grid.NewGrid[SpaceType](width, height)

	for y, line := range lines {
		for x, char := range line {
			var spaceType SpaceType

			switch char {
			case '.':
				spaceType = Empty
			case 'S':
				spaceType = Start
				start = geometry.Point{X: x, Y: y}
			case '^':
				spaceType = Splitter
			}
			_ = space.Set(&geometry.Point{X: x, Y: y}, spaceType)
		}
	}

	return
}
