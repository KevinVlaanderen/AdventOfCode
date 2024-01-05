package day25

import (
	"2023/src/framework"
	"github.com/samber/lo"
	"go/types"
	"golang.org/x/exp/slices"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/network"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/traverse"
	"hash/fnv"
	"regexp"
	"strings"
)

func Task1(data string, _ types.Nil) (result framework.Result[int]) {
	g := parse(data)

	nodes := g.Nodes()
	size := nodes.Len()
	nodes.Next()

	betweenness := lo.Entries(network.EdgeBetweenness(g))
	slices.SortFunc(betweenness, func(a, b lo.Entry[[2]int64, float64]) int {
		return int(b.Value - a.Value)
	})
	for _, entry := range betweenness[:3] {
		g.RemoveEdge(entry.Key[0], entry.Key[1])
	}

	found := 0
	search := traverse.BreadthFirst{
		Visit: func(node graph.Node) {
			found++
		},
	}
	search.Walk(g, nodes.Node(), nil)

	result.Value = found * (size - found)

	return
}

var wiringPattern = regexp.MustCompile(`(?m)^(.*): (.*)$`)

func parse(data string) *simple.UndirectedGraph {
	g := simple.NewUndirectedGraph()

	matches := wiringPattern.FindAllStringSubmatch(data, -1)
	for _, match := range matches {
		from := match[1]

		if node := g.Node(Component(from).ID()); node == nil {
			g.AddNode(Component(from))
		}
		for _, to := range strings.Split(match[2], " ") {
			if node := g.Node(Component(to).ID()); node == nil {
				g.AddNode(Component(to))
			}
			if edge := g.Edge(Component(from).ID(), Component(to).ID()); edge == nil {
				g.SetEdge(simple.Edge{F: Component(from), T: Component(to)})
			}
		}
	}
	return g
}

type Component string

func (c Component) ID() int64 {
	h := fnv.New64()
	_, _ = h.Write([]byte(c))
	return int64(h.Sum64() >> 1)
}
