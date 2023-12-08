package day8

import (
	"2023/src/framework/task"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"regexp"
	"strings"
)

func Task1(filePath string) (result task.Result[int]) {
	lines := task.ReadLines(filePath)
	steps := strings.Split(lines[0], "")
	network := NewNetwork(lines[2:])

	for i := 0; i <= len(steps); i++ {
		if network.current == network.target {
			break
		}
		if i == len(steps) {
			i = 0
		}

		network.Next(Direction(steps[i]))
		result.Value++
	}

	return
}

var nodePattern = regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

type Network struct {
	current, target int
	nodes           []lo.Tuple2[int, int]
}

func NewNetwork(lines []string) Network {
	nodeMatches := lo.Map(lines, func(line string, index int) []string {
		matches := nodePattern.FindAllStringSubmatch(line, -1)
		return matches[0]
	})

	nodeNames := lo.Map(nodeMatches, func(match []string, index int) string {
		return match[1]
	})

	return Network{
		current: lo.IndexOf(nodeNames, "AAA"),
		target:  lo.IndexOf(nodeNames, "ZZZ"),
		nodes: lop.Map(nodeMatches, func(match []string, index int) lo.Tuple2[int, int] {
			return lo.Tuple2[int, int]{
				A: lo.IndexOf(nodeNames, match[2]),
				B: lo.IndexOf(nodeNames, match[3]),
			}
		}),
	}
}

func (n *Network) Next(direction Direction) {
	switch direction {
	case Left:
		n.current = n.nodes[n.current].A
	case Right:
		n.current = n.nodes[n.current].B
	}
}

type Direction string

const (
	Left  Direction = "L"
	Right           = "R"
)
