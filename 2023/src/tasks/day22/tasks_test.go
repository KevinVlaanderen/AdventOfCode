package day22

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 5, Type: _testing.TestData},
			{Path: "day22", Expected: 424, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 7, Type: _testing.TestData},
			{Path: "day22", Expected: 55483, Type: _testing.RealData},
		},
	},
}

func TestDay22(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay22(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
