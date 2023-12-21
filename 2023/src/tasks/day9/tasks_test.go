package day9

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 114, Type: _testing.TestData},
			{Path: "day9", Expected: 1993300041, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 2, Type: _testing.TestData},
			{Path: "day9", Expected: 1038, Type: _testing.RealData},
		},
	},
}

func TestDay9(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay9(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
