package day4

import (
	_testing "2025/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 13, Type: _testing.TestData},
			{Path: "day4", Expected: 1491, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 43, Type: _testing.TestData},
			{Path: "day4", Expected: 8722, Type: _testing.RealData},
		},
	},
}

func TestDay1(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay1(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
