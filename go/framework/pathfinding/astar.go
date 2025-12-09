package pathfinding

import (
	"aoc/framework/geometry/geo2d"
	"errors"

	"github.com/oleiade/lane/v2"
	"golang.org/x/exp/slices"
)

type Graph interface {
	Neighbours(point PointWithDirection, mode geo2d.NeighbourMode) []geo2d.Point
	Cost(current PointWithDirection, next PointWithDirection) float64
	Heuristic(a, b geo2d.Point) float64
}

// type CostFn[G Graph]
// type HeuristicFn func

type AStar struct {
	graph     Graph
	cameFrom  map[geo2d.Point]PointWithDirection
	costSoFar map[PointWithDirection]float64
}

func NewAStar(graph Graph) *AStar {
	return &AStar{graph: graph}
}

type OrientationAware interface {
	OrientationOf(other geo2d.Point) (geo2d.Orientation, bool)
}

type PointWithDirection struct {
	Point     geo2d.Point
	Direction geo2d.Orientation
	Length    int
}

func (a *AStar) RunSearch(start geo2d.Point, goal geo2d.Point) {
	a.cameFrom = make(map[geo2d.Point]PointWithDirection)
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

		for _, next := range a.graph.Neighbours(current, geo2d.Orthogonal) {
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

func (a *AStar) CalculatePath(start, goal geo2d.Point) []geo2d.Point {
	path := make([]geo2d.Point, 0)
	for current := goal; current != start; current = a.cameFrom[current].Point {
		path = append(path, current)
	}
	path = append(path, start)
	slices.Reverse(path)
	return path
}
