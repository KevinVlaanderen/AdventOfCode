package day8

import (
	math2 "aoc/framework/geometry/geo3d"
	"aoc/framework/tasks"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/samber/lo"
)

type Connection struct {
	From, To math2.Point3D
}

func hash(p math2.Point3D) string {
	return fmt.Sprintf("%v,%v,%v", p.X, p.Y, p.Z)
}

type Distance struct {
	Key   Connection
	Value float64
}

func Task1(data string, nConnections int) (result tasks.Result[int]) {
	boxes := parse(data)

	distances := calculateDistances(boxes)

	g := createGraph(boxes, distances, nConnections)

	clusters := clusterPoints(g)

	result.Value = lo.Reduce(clusters[:3], func(agg int, cluster []math2.Point3D, index int) int {
		return agg * len(cluster)
	}, 1)

	return
}

func Task2(data string, nConnections int) (result tasks.Result[int]) {
	boxes := parse(data)

	distances := calculateDistances(boxes)

	g := createGraph(boxes, distances, nConnections)

	index := nConnections - 1
	done := false
	for !done {
		index++
		distance := distances[index]
		_ = g.AddEdge(hash(distance.Key.From), hash(distance.Key.To))

		adjacencyMap, _ := g.AdjacencyMap()
		if lo.EveryBy(lo.Keys(adjacencyMap), func(item string) bool {
			return len(adjacencyMap[item]) > 0
		}) {
			done = true
		}
	}

	finalDistance := distances[index]

	result.Value = int(finalDistance.Key.From.X) * int(finalDistance.Key.To.X)

	return
}

func parse(data string) []math2.Point3D {
	lines := tasks.Lines(data)

	return lo.Map(lines, func(line string, index int) math2.Point3D {
		parts := strings.Split(line, ",")
		partsInt := lo.Map(parts, func(part string, index int) int {
			n, _ := strconv.Atoi(part)
			return n
		})
		return math2.Point3D{X: float64(partsInt[0]), Y: float64(partsInt[1]), Z: float64(partsInt[2])}
	})
}

func createGraph(points []math2.Point3D, distances []Distance, n int) graph.Graph[string, math2.Point3D] {
	g := graph.New(hash, graph.Acyclic())

	for _, point := range points {
		_ = g.AddVertex(point)
	}

	nDistances := distances
	if n > 0 {
		nDistances = distances[:n]
	}

	for _, kv := range nDistances {
		_ = g.AddEdge(hash(kv.Key.From), hash(kv.Key.To))
	}

	return g
}

func calculateDistances(points []math2.Point3D) (distances []Distance) {
	distanceMap := make(map[Connection]float64)

	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			distance := math2.Distance3D(points[i], points[j])
			distanceMap[Connection{From: points[i], To: points[j]}] = distance
		}
	}

	for k, v := range distanceMap {
		distances = append(distances, Distance{k, v})
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Value < distances[j].Value
	})

	return
}

func clusterPoints(g graph.Graph[string, math2.Point3D]) [][]math2.Point3D {
	clusters := make([][]math2.Point3D, 0)
	adjacencyMap, _ := g.AdjacencyMap()

	for src := range adjacencyMap {
		srcVertex, _ := g.Vertex(src)

		if lo.EveryBy(clusters, func(cluster []math2.Point3D) bool {
			return !lo.Contains(cluster, srcVertex)
		}) {
			cluster := []math2.Point3D{srcVertex}

			_ = graph.DFS(g, src, func(s string) bool {
				v, _ := g.Vertex(s)
				if v != srcVertex {
					cluster = append(cluster, v)
				}
				return false
			})

			clusters = append(clusters, cluster)
		}
	}

	sort.Slice(clusters, func(i, j int) bool {
		return len(clusters[i]) > len(clusters[j])
	})
	return clusters
}
