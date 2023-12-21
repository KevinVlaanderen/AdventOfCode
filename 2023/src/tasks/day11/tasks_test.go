package day11

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 374, Type: _testing.TestData},
			{Path: "day11", Expected: 9418609, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "day11", Expected: 593821230983, Type: _testing.RealData},
		},
	},
}

func TestDay11(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay11(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
