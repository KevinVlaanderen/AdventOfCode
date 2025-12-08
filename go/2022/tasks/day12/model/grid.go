package model

import (
	"aoc/framework/math"
	"errors"

	"github.com/oleiade/lane/v2"
)

type Grid struct {
	Squares     [][]int
	Start       Position
	Destination Position
}

type Position struct {
	X int
	Y int
}

func (g *Grid) PathTo(from Position, to Position) (cameFrom map[Position]Position, err error) {
	frontier := lane.NewPriorityQueue[Position, float64](func(lhs, rhs float64) bool {
		return lhs < rhs
	})

	frontier.Push(from, 1)

	cameFrom = map[Position]Position{}
	costSoFar := map[Position]int{}
	// cameFrom[g.Start] = nil
	costSoFar[from] = 0

	for !frontier.Empty() {
		current, _, _ := frontier.Pop()

		if current == to {
			break
		}

		for _, next := range g.neighbours(current) {
			newCost := costSoFar[current] + g.cost(current, next)
			if _, ok := costSoFar[next]; !ok || newCost < costSoFar[next] {
				costSoFar[next] = newCost
				priority := 1 / float64(newCost+g.heuristic(to, next))
				frontier.Push(next, priority)

				cameFrom[next] = current
			}
		}
	}

	if _, found := cameFrom[to]; !found {
		return nil, errors.New("no path found")
	}

	return cameFrom, nil
}

func (g *Grid) neighbours(position Position) (neighbours []Position) {
	if position.X >= 1 && g.Squares[position.X-1][position.Y] <= g.Squares[position.X][position.Y]+1 {
		neighbours = append(neighbours, Position{X: position.X - 1, Y: position.Y})
	}
	if position.X < len(g.Squares)-1 && g.Squares[position.X+1][position.Y] <= g.Squares[position.X][position.Y]+1 {
		neighbours = append(neighbours, Position{X: position.X + 1, Y: position.Y})
	}
	if position.Y >= 1 && g.Squares[position.X][position.Y-1] <= g.Squares[position.X][position.Y]+1 {
		neighbours = append(neighbours, Position{X: position.X, Y: position.Y - 1})
	}
	if position.Y < len(g.Squares[0])-1 && g.Squares[position.X][position.Y+1] <= g.Squares[position.X][position.Y]+1 {
		neighbours = append(neighbours, Position{X: position.X, Y: position.Y + 1})
	}
	return
}

func (g *Grid) cost(current Position, next Position) int {
	return 1
}

func (g *Grid) heuristic(a Position, b Position) int {
	return math.AbsInt(a.X-b.X) + math.AbsInt(a.Y-b.Y)
}
