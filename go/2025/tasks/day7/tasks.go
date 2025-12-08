package day7

import (
	"aoc/framework"
	"aoc/framework/geometry"
	"aoc/framework/geometry/grid"
	"fmt"
	"go/types"

	"github.com/dominikbraun/graph"
	"github.com/oleiade/lane/v2"
)

type SpaceType uint8

const (
	Start SpaceType = iota
	Empty
	Splitter
	Beam
	End
)

type Vertex struct {
	Type     SpaceType
	Position geometry.Point
}

type Edge struct {
	From      Vertex
	Direction geometry.Direction
	Position  geometry.Point
}

func hash(vertex Vertex) string {
	return fmt.Sprintf("%v-%v-%v", vertex.Position.X, vertex.Position.Y, vertex.Type)
}

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

func Task2(data string, _ types.Nil) (result framework.Result[int]) {
	lab, start := parse(data)

	g, startVertex, endVertex := createGraph(lab, start)

	m, _ := g.AdjacencyMap()
	cache := make(map[string]int)
	result.Value = dfs(m, hash(startVertex), hash(endVertex), graph.Edge[string]{}, cache)

	return
}

func createGraph(lab grid.Grid[SpaceType], start geometry.Point) (graph.Graph[string, Vertex], Vertex, Vertex) {
	g := graph.New(hash, graph.Directed(), graph.Acyclic())

	startVertex := Vertex{Type: Start, Position: start}
	_ = g.AddVertex(startVertex)

	endVertex := Vertex{Type: End}
	_ = g.AddVertex(endVertex)

	beams := lane.NewQueue[*Edge](&Edge{From: startVertex, Position: start.Neighbour(geometry.South)})

	for beams.Size() > 0 {
		currentBeam, _ := beams.Head()
		nextPosition := currentBeam.Position.Neighbour(geometry.South)
		spaceType, found := lab.Get(&nextPosition)
		if found {
			if spaceType == Empty {
				currentBeam.Position = nextPosition
			} else if spaceType == Splitter {
				existingVertex, err := g.Vertex(hash(Vertex{Type: Splitter, Position: nextPosition}))

				if err == nil {
					_ = g.AddEdge(hash(currentBeam.From), hash(existingVertex), graph.EdgeData(1))
					beams.Dequeue()
				} else {
					newVertex := Vertex{Type: Splitter, Position: nextPosition}
					_ = g.AddVertex(newVertex)
					_ = g.AddEdge(hash(currentBeam.From), hash(newVertex), graph.EdgeData(1))

					left := nextPosition.Neighbour(geometry.West)
					beams.Enqueue(&Edge{From: newVertex, Direction: geometry.Left, Position: left})

					right := nextPosition.Neighbour(geometry.East)
					beams.Enqueue(&Edge{From: newVertex, Direction: geometry.Right, Position: right})

					beams.Dequeue()
				}
			}
		} else {
			if edge, err := g.Edge(hash(currentBeam.From), hash(endVertex)); err != nil {
				_ = g.AddEdge(hash(currentBeam.From), hash(endVertex), graph.EdgeData(1))
			} else {
				_ = g.UpdateEdge(hash(currentBeam.From), hash(endVertex), graph.EdgeData(edge.Properties.Data.(int)+1))
			}

			beams.Dequeue()
		}
	}
	return g, startVertex, endVertex
}

func dfs(m map[string]map[string]graph.Edge[string], src string, dest string, edge graph.Edge[string], cache map[string]int) int {
	if src == dest {
		return edge.Properties.Data.(int)
	} else if count, ok := cache[src]; ok {
		return count
	} else {
		for adj, e := range m[src] {
			count += dfs(m, adj, dest, e, cache)
		}
		cache[src] = count
		return count
	}
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
