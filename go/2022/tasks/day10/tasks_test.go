package day10

import (
	_testing "aoc/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions1 = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 13140, Type: _testing.TestData},
			{Path: "day10", Expected: 12740, Type: _testing.RealData},
		},
	},
}

var taskDefinitions2 = []_testing.TaskDefinition[string, types.Nil]{
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[string, types.Nil]{
			{Path: "day10", Expected: "RBPARAGF", Type: _testing.RealData},
		},
	},
}

func Test1(t *testing.T) {
	_testing.RunTests(t, taskDefinitions1)
}

func Test2(t *testing.T) {
	_testing.RunTests(t, taskDefinitions2)
}

func Benchmark1(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions1)
}

func Benchmark2(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions2)
}
