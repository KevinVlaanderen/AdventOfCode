package day3

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 4361, Type: _testing.TestData},
			{Path: "day3", Expected: 546312, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 467835, Type: _testing.TestData},
			{Path: "day3", Expected: 87449461, Type: _testing.RealData},
		},
	},
}

func TestDay3(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay3(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
