package day6

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 288, Type: _testing.TestData},
			{Path: "day6", Expected: 252000, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 71503, Type: _testing.TestData},
			{Path: "day6", Expected: 36992486, Type: _testing.RealData},
		},
	},
}

func TestDay6(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay6(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
