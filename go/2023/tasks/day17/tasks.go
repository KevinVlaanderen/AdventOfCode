package day17

import (
	"aoc/framework/geometry/geo2d"
	"aoc/framework/pathfinding"
	"aoc/framework/tasks"
	"go/types"
	"math"
	"strconv"

	"gonum.org/v1/gonum/graph/simple"
)

type Graph struct {
	*simple.WeightedDirectedGraph
}

func (g *Graph) Cost(current pathfinding.PointWithDirection, next pathfinding.PointWithDirection) float64 {
	return g.WeightedEdge(current.Point.ID(), next.Point.ID()).Weight()
}

func (g *Graph) Heuristic(from, to geo2d.Point) float64 {
	return math.Abs(float64(from.X-to.X)) + math.Abs(float64(from.Y-to.Y))
}

func (g *Graph) Neighbours(point pathfinding.PointWithDirection, _ geo2d.NeighbourMode) []geo2d.Point {
	nodes := g.From(point.Point.ID())
	points := make([]geo2d.Point, 0, nodes.Len())
	for nodes.Next() {
		nextPoint := nodes.Node().(geo2d.Point)
		direction, _ := point.Point.OrientationOf(nextPoint)
		if point.Point != nextPoint && direction != geo2d.OppositeOrientation[point.Direction] {
			points = append(points, nextPoint)
		}
	}
	return points
}

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	g, width, height := parse(data)

	graph := &Graph{g}

	start := geo2d.Point{X: 0, Y: 0}
	goal := geo2d.Point{X: int(width - 1), Y: int(height - 1)}

	aStar := pathfinding.NewAStar(graph)
	aStar.RunSearch(start, goal)
	path := aStar.CalculatePath(start, goal)

	println(path)

	// dStar := dynamic.NewDStarLite(start, goal, g, func(a, b graph.Node) float64 {
	//	xA, xB, yA, yB := a.ID()%width, b.ID()%width, a.ID()/width, b.ID()/width
	//	return math.Abs(float64(xA-xB)) + math.Abs(float64(yA-yB))
	// }, simple.NewWeightedDirectedGraph(0, math.Inf(1)))
	//
	// currentPoint := start
	// var currentHeading geometry.Orientation
	// numInHeading := -1
	// var oldEdge graph.Edge
	// for dStar.Step() {
	//	//if oldEdge != nil {
	//	//	dStar.UpdateWorld([]graph.Edge{oldEdge})
	//	//	oldEdge = nil
	//	//}
	//
	//	nextPoint := dStar.Here().(geometry.Point)
	//	heading, _ := currentPoint.OrientationOf(nextPoint)
	//
	//	if numInHeading == -1 || heading != currentHeading {
	//		currentHeading = heading
	//		numInHeading = 1
	//	} else {
	//		numInHeading++
	//	}
	//
	//	if numInHeading == 3 {
	//		previousOldEdge := oldEdge
	//		if oldEdge = g.Edge(nextPoint.ID(), nextPoint.Neighbour(currentHeading).ID()); oldEdge != nil {
	//			newEdge := simple.WeightedEdge{F: oldEdge.U(), T: oldEdge.V(), W: 999999999999}
	//			changes := []graph.Edge{newEdge}
	//			if previousOldEdge != nil {
	//				changes = append(changes, previousOldEdge)
	//			}
	//			dStar.UpdateWorld(changes)
	//		}
	//	}
	//	currentPoint = nextPoint
	// }
	//
	// dStar.MoveTo(start)
	// p, weight := dStar.Path()
	//
	// fmt.Print(p, weight)

	// aStar := pathfinding.NewAStar[string, geometry.Point](g)
	// aStar.RunSearch(start, goal, aStarCost, heuristic, validPath)
	// path := aStar.CalculatePath(start, goal)

	// paths := YenKSP(g, start, goal, 10, nil)
	// println(len(path))

	// result.Value = lo.Sum(lo.Map(path, func(hash string, index int) int {
	//	_, props, _ := g.VertexWithProperties(hash)
	//	return props.Weight
	// }))

	return
}

type Node struct {
	id int64
}

func (n Node) ID() int64 {
	return n.id
}

func parse(data string) (*simple.WeightedDirectedGraph, int64, int64) {
	lines := tasks.Lines(data)

	width, height := len(lines[0]), len(lines)

	g := simple.NewWeightedDirectedGraph(0, math.Inf(1))

	costs := make([][]float64, height)
	for i := 0; i < height; i++ {
		costs[i] = make([]float64, width)
	}

	for y, row := range lines {
		for x, value := range row {
			var cost float64

			if number, err := strconv.Atoi(string(value)); err != nil {
				panic(err)
			} else {
				cost = float64(number)
			}

			costs[y][x] = cost

			point := geo2d.Point{X: x, Y: y}

			g.AddNode(point)

			if x > 0 {
				westPoint := point.Neighbour(geo2d.West)
				edge := simple.WeightedEdge{F: point, T: westPoint, W: costs[y][x-1]}
				g.SetWeightedEdge(edge)
				edge = simple.WeightedEdge{F: westPoint, T: point, W: cost}
				g.SetWeightedEdge(edge)
			}

			if y > 0 {
				northPoint := point.Neighbour(geo2d.North)
				edge := simple.WeightedEdge{F: point, T: northPoint, W: costs[y-1][x]}
				g.SetWeightedEdge(edge)
				edge = simple.WeightedEdge{F: northPoint, T: point, W: cost}
				g.SetWeightedEdge(edge)
			}
		}

	}
	return g, int64(width), int64(height)
}
