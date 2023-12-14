package model

import (
	"2023/src/framework/geometry"
	"github.com/samber/lo"
	"github.com/samber/lo/parallel"
	"regexp"
	"strings"
)

type Network struct {
	steps     []geometry.Direction
	nodeNames []string
	nodes     []lo.Tuple2[int, int]
}

func NewNetwork(stepData string, nodeData string) Network {
	steps := parseSteps(stepData)
	nodes := parseNodes(nodeData)

	nodeNames := lo.Map(nodes, func(item lo.Tuple3[string, string, string], index int) string {
		return item.A
	})

	return Network{
		steps:     steps,
		nodeNames: nodeNames,
		nodes: parallel.Map(nodes, func(item lo.Tuple3[string, string, string], index int) lo.Tuple2[int, int] {
			return lo.Tuple2[int, int]{
				A: lo.IndexOf(nodeNames, item.B),
				B: lo.IndexOf(nodeNames, item.C),
			}
		}),
	}
}

func parseSteps(line string) []geometry.Direction {
	return lo.Map(strings.Split(line, ""), func(direction string, index int) geometry.Direction {
		switch direction {
		case "L":
			return geometry.Left
		case "R":
			return geometry.Right
		default:
			panic("unknown direction")
		}
	})
}

var nodePattern = regexp.MustCompile(`(?m)^(\w+) = \((\w+), (\w+)\)$`)

func parseNodes(data string) []lo.Tuple3[string, string, string] {
	matches := nodePattern.FindAllStringSubmatch(data, -1)

	return lo.Map(matches, func(match []string, index int) lo.Tuple3[string, string, string] {
		return lo.Tuple3[string, string, string]{A: match[1], B: match[2], C: match[3]}
	})
}

func (n Network) IndicesBy(filter func(item string, index int) bool) []int {
	return lo.Map(lo.Filter(n.nodeNames, filter), func(name string, index int) int {
		return n.IndexOf(name)
	})
}

func (n Network) Next(current int, direction geometry.Direction) int {
	switch direction {
	case geometry.Left:
		return n.nodes[current].A
	case geometry.Right:
		return n.nodes[current].B
	}
	panic("unknown direction")
}

func (n Network) FindRequiredStepsBy(start int, fn func(item int) bool) (result int) {
	current := start
	numSteps := len(n.steps)
	for i := 0; i <= numSteps; i++ {
		if fn(current) {
			break
		}
		if i == numSteps {
			i = 0
		}

		current = n.Next(current, n.steps[i])
		result++
	}
	return
}

func (n Network) NameOf(index int) string {
	return n.nodeNames[index]
}

func (n Network) IndexOf(name string) int {
	return lo.IndexOf(n.nodeNames, name)
}
