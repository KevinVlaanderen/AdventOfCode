package model

import (
	"2022/src/framework"
	"container/heap"
	"errors"
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

func ReadGrid(data []string) *Grid {
	grid := Grid{}
	grid.Squares = make([][]int, len(data[0]))
	for i := range grid.Squares {
		grid.Squares[i] = make([]int, len(data))
	}

	for y := range data {
		for x, char := range []rune(data[y]) {
			switch char {
			case 'S':
				grid.Start = Position{X: x, Y: y}
				grid.Squares[x][y] = 1
			case 'E':
				grid.Destination = Position{X: x, Y: y}
				grid.Squares[x][y] = 26
			default:
				grid.Squares[x][y] = int(char) - 96
			}
		}
	}

	return &grid
}

func (g *Grid) PathTo(from Position, to Position) (cameFrom map[Position]Position, err error) {
	frontier := make(framework.PriorityQueue[Position], 1)

	frontier[0] = &framework.Item[Position]{
		Value:    from,
		Priority: 1,
		Index:    0,
	}

	heap.Init(&frontier)

	cameFrom = map[Position]Position{}
	costSoFar := map[Position]int{}
	//cameFrom[g.Start] = nil
	costSoFar[from] = 0

	for frontier.Len() > 0 {
		current := heap.Pop(&frontier).(*framework.Item[Position])

		if current.Value == to {
			break
		}

		for _, next := range g.neighbours(current.Value) {
			newCost := costSoFar[current.Value] + g.cost(current.Value, next)
			if _, ok := costSoFar[next]; !ok || newCost < costSoFar[next] {
				costSoFar[next] = newCost
				priority := 1 / float64(newCost+g.heuristic(to, next))
				item := &framework.Item[Position]{
					Value:    next,
					Priority: priority,
				}
				heap.Push(&frontier, item)
				frontier.Update(item, item.Value, priority)

				cameFrom[next] = current.Value
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
	return framework.AbsInt(a.X-b.X) + framework.AbsInt(a.Y-b.Y)
}
