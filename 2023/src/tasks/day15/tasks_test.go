package day15

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 52, Type: _testing.TestData},
			{Path: "data2", Expected: 1320, Type: _testing.TestData},
			{Path: "day15", Expected: 505427, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data2", Expected: 145, Type: _testing.TestData},
			{Path: "day15", Expected: 243747, Type: _testing.RealData},
		},
	},
}

func TestDay15(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay15(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
