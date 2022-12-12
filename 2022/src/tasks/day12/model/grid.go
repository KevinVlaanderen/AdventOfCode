package model

import (
	"2022/src/framework"
	"container/heap"
)

type Grid struct {
	squares     [][]int
	Start       Position
	Destination Position
}

type Position struct {
	x int
	y int
}

func ReadGrid(data []string) *Grid {
	grid := Grid{}
	grid.squares = make([][]int, len(data[0]))
	for i := range grid.squares {
		grid.squares[i] = make([]int, len(data))
	}

	for y := range data {
		for x, char := range []rune(data[y]) {
			switch char {
			case 'S':
				grid.Start = Position{x: x, y: y}
				grid.squares[x][y] = 1
			case 'E':
				grid.Destination = Position{x: x, y: y}
				grid.squares[x][y] = 26
			default:
				grid.squares[x][y] = int(char) - 96
			}
		}
	}

	return &grid
}

func (g *Grid) PathTo(position Position) map[Position]Position {
	frontier := make(framework.PriorityQueue[Position], 1)

	frontier[0] = &framework.Item[Position]{
		Value:    g.Start,
		Priority: 1,
		Index:    0,
	}

	heap.Init(&frontier)

	cameFrom := map[Position]Position{}
	costSoFar := map[Position]int{}
	//cameFrom[g.Start] = nil
	costSoFar[g.Start] = 0

	for frontier.Len() > 0 {
		current := frontier.Pop().(*framework.Item[Position])

		if current.Value == position {
			break
		}

		for _, next := range g.neighbours(current.Value) {
			newCost := costSoFar[current.Value] + g.cost(current.Value, next)
			if _, ok := costSoFar[next]; !ok || newCost < costSoFar[next] {
				costSoFar[next] = newCost
				priority := newCost + g.heuristic(g.Destination, next)
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

	return cameFrom
}

func (g *Grid) neighbours(position Position) (neighbours []Position) {
	if position.x >= 1 && g.squares[position.x-1][position.y] <= g.squares[position.x][position.y]+1 {
		neighbours = append(neighbours, Position{x: position.x - 1, y: position.y})
	}
	if position.x < len(g.squares)-1 && g.squares[position.x+1][position.y] <= g.squares[position.x][position.y]+1 {
		neighbours = append(neighbours, Position{x: position.x + 1, y: position.y})
	}
	if position.y >= 1 && g.squares[position.x][position.y-1] <= g.squares[position.x][position.y]+1 {
		neighbours = append(neighbours, Position{x: position.x, y: position.y - 1})
	}
	if position.y < len(g.squares[0])-1 && g.squares[position.x][position.y+1] <= g.squares[position.x][position.y]+1 {
		neighbours = append(neighbours, Position{x: position.x, y: position.y + 1})
	}
	return
}

func (g *Grid) cost(current Position, next Position) int {
	return 1
}

func (g *Grid) heuristic(a Position, b Position) int {
	return framework.AbsInt(a.x-b.x) + framework.AbsInt(a.y-b.y)
}
