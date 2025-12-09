package day23

import (
	_testing "aoc/framework/tasks/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 94, Type: _testing.TestData},
			{Path: "day23", Expected: 2034, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 154, Type: _testing.TestData},
			{Path: "day23", Expected: 6302, Type: _testing.RealData},
		},
	},
}

func Test(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
