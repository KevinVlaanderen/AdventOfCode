package day14

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 136, Type: _testing.TestData},
			{Path: "day14", Expected: 113525, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 64, Type: _testing.TestData},
			{Path: "day14", Expected: 101292, Type: _testing.RealData},
		},
	},
}

func TestDay14(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay14(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
