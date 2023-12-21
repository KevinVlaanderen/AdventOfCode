package day18

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 62, Type: _testing.TestData},
			{Path: "day18", Expected: 70026, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 952408144115, Type: _testing.TestData},
			{Path: "day18", Expected: 68548301037382, Type: _testing.RealData},
		},
	},
}

func TestDay18(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay18(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
