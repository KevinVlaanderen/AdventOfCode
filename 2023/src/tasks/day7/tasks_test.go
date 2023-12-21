package day7

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 6440, Type: _testing.TestData},
			{Path: "day7", Expected: 251216224, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 5905, Type: _testing.TestData},
			{Path: "day7", Expected: 250825971, Type: _testing.RealData},
		},
	},
}

func TestDay7(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay7(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
