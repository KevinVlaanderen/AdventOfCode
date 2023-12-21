package day13

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 405, Type: _testing.TestData},
			{Path: "day13", Expected: 30802, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 400, Type: _testing.TestData},
			{Path: "day13", Expected: 37876, Type: _testing.RealData},
		},
	},
}

func TestDay13(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay13(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
