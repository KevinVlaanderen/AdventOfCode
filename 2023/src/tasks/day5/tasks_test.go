package day5

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 35, Type: _testing.TestData},
			{Path: "day5", Expected: 484023871, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 46, Type: _testing.TestData},
			{Path: "day5", Expected: 46294175, Type: _testing.RealData},
		},
	},
}

func TestDay5(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay5(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
