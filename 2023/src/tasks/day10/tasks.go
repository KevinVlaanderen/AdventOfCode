package day10

import (
	"2023/src/framework"
	"2023/src/framework/geometry"
)

func Task1(filePath string) (result framework.Result[int]) {
	data := framework.ReadAll(filePath)
	surface := NewSurface(data)

	start := surface.start
	loop := []geometry.Point{surface.start}

	current, previous := surface.start, geometry.Point{X: -1, Y: -1}
	for {
		neighbour := <-surface.grid.NeighboursBy(current, func(neighbour geometry.Point) bool {
			var side Direction
			switch {
			case neighbour.Y < current.Y:
				side = Top
			case neighbour.X > current.X:
				side = Right
			case neighbour.Y > current.Y:
				side = Bottom
			case neighbour.X < current.X:
				side = Left
			}
			return (current.X == neighbour.X || current.Y == neighbour.Y) &&
				surface.grid[current].ConnectsTo(surface.grid[neighbour], side) &&
				neighbour != previous
		})
		if neighbour == start {
			break
		} else {
			loop = append(loop, neighbour)
			current, previous = neighbour, current
		}
	}

	result.Value = len(loop) / 2

	return
}

type Surface struct {
	start geometry.Point
	grid  geometry.Grid[Pipe]
}

func NewSurface(data string) Surface {
	surface := Surface{
		grid: geometry.Grid[Pipe]{},
	}

	current := geometry.Point{}
	for _, r := range data {
		switch r {
		case '\n':
			current.X = 0
			current.Y++
			continue
		case 'S':
			surface.start = current
			surface.grid[current] = NewPipe(r)
		default:
			surface.grid[current] = NewPipe(r)
		}
		current.X++
	}
	return surface
}

type Pipe struct {
	Top, Right, Bottom, Left bool
}

func NewPipe(data rune) Pipe {
	switch data {
	case '|':
		return Pipe{Top: true, Bottom: true}
	case '-':
		return Pipe{Left: true, Right: true}
	case 'F':
		return Pipe{Right: true, Bottom: true}
	case 'J':
		return Pipe{Top: true, Left: true}
	case '7':
		return Pipe{Left: true, Bottom: true}
	case 'L':
		return Pipe{Top: true, Right: true}
	case '.':
		return Pipe{}
	case 'S':
		return Pipe{Top: true, Right: true, Bottom: true, Left: true}
	default:
		panic("unknown character")
	}
}

func (p Pipe) ConnectsTo(other Pipe, side Direction) bool {
	if (side == Top && !p.Top) ||
		(side == Right && !p.Right) ||
		(side == Bottom && !p.Bottom) ||
		(side == Left && !p.Left) {
		return false
	}
	oppositeSide := OppositeDirection[side]
	if (oppositeSide == Top && !other.Top) ||
		(oppositeSide == Right && !other.Right) ||
		(oppositeSide == Bottom && !other.Bottom) ||
		(oppositeSide == Left && !other.Left) {
		return false
	}
	return true
}

type Direction int

const (
	Top Direction = iota
	Right
	Bottom
	Left
)

var OppositeDirection = map[Direction]Direction{Top: Bottom, Right: Left, Bottom: Top, Left: Right}
