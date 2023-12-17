package pathfinding

import (
	"2023/src/framework/geometry"
	"errors"
	"github.com/oleiade/lane/v2"
	"golang.org/x/exp/slices"
)

type Graph interface {
	Neighbours(point PointWithDirection, mode geometry.NeighbourMode) []geometry.Point
	Cost(current PointWithDirection, next PointWithDirection) float64
	Heuristic(a, b geometry.Point) float64
}

//type CostFn[G Graph]
//type HeuristicFn func

type AStar struct {
	graph     Graph
	cameFrom  map[geometry.Point]PointWithDirection
	costSoFar map[PointWithDirection]float64
}

func NewAStar(graph Graph) *AStar {
	return &AStar{graph: graph}
}

type OrientationAware interface {
	OrientationOf(other geometry.Point) (geometry.Orientation, bool)
}

type PointWithDirection struct {
	Point     geometry.Point
	Direction geometry.Orientation
	Length    int
}

func (a *AStar) RunSearch(start geometry.Point, goal geometry.Point) {
	a.cameFrom = make(map[geometry.Point]PointWithDirection)
	a.costSoFar = make(map[PointWithDirection]float64)

	startWithDirection := PointWithDirection{Point: start}

	frontier := lane.NewMinPriorityQueue[PointWithDirection, float64]()
	frontier.Push(startWithDirection, 0)

	a.cameFrom[start] = startWithDirection
	a.costSoFar[startWithDirection] = 0

	for frontier.Size() > 0 {
		current, _, ok := frontier.Pop()
		if !ok {
			panic(errors.New("failed to dequeue"))
		}

		if current.Point == goal {
			break
		}

		for _, next := range a.graph.Neighbours(current, geometry.Orthogonal) {
			direction, _ := current.Point.OrientationOf(next)
			nextWithDirection := PointWithDirection{Point: next, Direction: direction}
			if current.Length == 0 || current.Direction != direction {
				nextWithDirection.Direction = direction
				nextWithDirection.Length = 1
			} else {
				nextWithDirection.Length++
			}

			if nextWithDirection.Length > 3 {
				continue
			}

			newCost := a.costSoFar[current] + a.graph.Cost(current, nextWithDirection)

			if costSoFar, found := a.costSoFar[nextWithDirection]; !found || newCost < costSoFar {
				a.costSoFar[nextWithDirection] = newCost
				priority := newCost + a.graph.Heuristic(next, goal)
				frontier.Push(nextWithDirection, priority)
				a.cameFrom[next] = current
			}
		}
	}
}

func (a *AStar) CalculatePath(start, goal geometry.Point) []geometry.Point {
	path := make([]geometry.Point, 0)
	for current := goal; current != start; current = a.cameFrom[current].Point {
		path = append(path, current)
	}
	path = append(path, start)
	slices.Reverse(path)
	return path
}
