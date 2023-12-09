package model

import (
	"github.com/samber/lo"
	"github.com/samber/lo/parallel"
	"regexp"
	"strings"
)

var nodePattern = regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

type Network struct {
	steps     []Direction
	nodeNames []string
	nodes     []lo.Tuple2[int, int]
}

func NewNetwork(stepData string, nodeData []string) Network {
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

func parseSteps(line string) []Direction {
	return lo.Map(strings.Split(line, ""), func(direction string, index int) Direction {
		switch direction {
		case "L":
			return Left
		case "R":
			return Right
		default:
			panic("unknown direction")
		}
	})
}

func parseNodes(lines []string) []lo.Tuple3[string, string, string] {
	return lo.Map(lines, func(line string, index int) lo.Tuple3[string, string, string] {
		match := nodePattern.FindAllStringSubmatch(line, -1)[0]
		return lo.Tuple3[string, string, string]{A: match[1], B: match[2], C: match[3]}
	})
}

func (n Network) IndicesBy(filter func(item string, index int) bool) []int {
	return lo.Map(lo.Filter(n.nodeNames, filter), func(name string, index int) int {
		return n.IndexOf(name)
	})
}

func (n Network) Next(current int, direction Direction) int {
	switch direction {
	case Left:
		return n.nodes[current].A
	case Right:
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