package day16

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 46, Type: _testing.TestData},
			{Path: "day16", Expected: 7210, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 51, Type: _testing.TestData},
			{Path: "day16", Expected: 7673, Type: _testing.RealData},
		},
	},
}

func TestDay16(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay16(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
