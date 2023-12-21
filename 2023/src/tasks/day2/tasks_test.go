package day2

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 8, Type: _testing.TestData},
			{Path: "day2", Expected: 2278, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 2286, Type: _testing.TestData},
			{Path: "day2", Expected: 67953, Type: _testing.RealData},
		},
	},
}

func TestDay2(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay2(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
