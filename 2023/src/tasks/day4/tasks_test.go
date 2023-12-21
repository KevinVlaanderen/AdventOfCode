package day4

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 13, Type: _testing.TestData},
			{Path: "day4", Expected: 22193, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 30, Type: _testing.TestData},
			{Path: "day4", Expected: 5625994, Type: _testing.RealData},
		},
	},
}

func TestDay4(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay4(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
