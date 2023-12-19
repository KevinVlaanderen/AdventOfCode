package model

import (
	"fmt"
	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"github.com/samber/lo"
	"os"
	"path"
)

type System struct {
	graph graph.Graph[string, string]
}

func NewSystem(workflows map[string][]Rule) System {
	g := graph.New(graph.StringHash, graph.Rooted(), graph.Directed())
	for name, rules := range workflows {
		_ = g.AddVertex(name)

		allSame := true
		for i := 0; i < len(rules)-1; i++ {
			if rules[i].Next != rules[i+1].Next {
				allSame = false
				break
			}
		}
		if allSame {
			_ = g.AddEdge(name, rules[0].Next, graph.EdgeData(nil))
			continue
		}

		for ruleIndex, rule := range rules {
			_ = g.AddVertex(rule.Next)

			currentSource := name
			currentTarget := rule.Next
			currentCondition := rule.Condition

			previousRules := rules[:ruleIndex]

			for previousRuleIndex, previousRule := range previousRules {
				currentTarget = fmt.Sprintf("%v_%v_%v", name, ruleIndex, previousRuleIndex)
				_ = g.AddVertex(currentTarget)
				_ = g.AddEdge(currentSource, currentTarget, graph.EdgeData(previousRule.Condition.Opposite()))
				currentSource = currentTarget
			}

			currentTarget = rule.Next

			_ = g.AddEdge(currentSource, currentTarget, graph.EdgeData(currentCondition))
		}
	}

	return System{g}
}

func (s *System) ValidBoundaries(start, end string) []map[Category]lo.Tuple2[int, int] {
	validPaths, _ := graph.AllPathsBetween(s.graph, start, end)
	validBoundaries := make([]map[Category]lo.Tuple2[int, int], len(validPaths))

	for index, p := range validPaths {
		boundaries := map[Category]lo.Tuple2[int, int]{
			"x": {A: 1, B: 4000},
			"m": {A: 1, B: 4000},
			"a": {A: 1, B: 4000},
			"s": {A: 1, B: 4000},
		}

		for i := 0; i < len(p)-1; i++ {
			edge, _ := s.graph.Edge(p[i], p[i+1])

			if edge.Properties.Data == nil {
				continue
			}

			condition := edge.Properties.Data.(*Condition)
			if condition == nil {
				continue
			}

			switch condition.Operator {
			case LessThan:
				if boundaries[condition.Category].B >= condition.N {
					newBoundary := boundaries[condition.Category]
					newBoundary.B = condition.N - 1
					boundaries[condition.Category] = newBoundary
				}
			case GreaterThan:
				if boundaries[condition.Category].A <= condition.N {
					newBoundary := boundaries[condition.Category]
					newBoundary.A = condition.N + 1
					boundaries[condition.Category] = newBoundary
				}
			}
		}
		validBoundaries[index] = boundaries
	}
	return validBoundaries
}

func (s *System) Draw(filename string) {
	edges, _ := s.graph.Edges()

	for _, edge := range edges {
		label := "always"
		if edge.Properties.Data != nil {
			condition := edge.Properties.Data.(*Condition)
			if condition != nil {
				label = condition.String()
			}
		}

		_ = s.graph.UpdateEdge(edge.Source, edge.Target, graph.EdgeAttribute("label", label))
	}

	file, _ := os.Create(path.Join(".", filename))
	_ = draw.DOT(s.graph, file)
}
