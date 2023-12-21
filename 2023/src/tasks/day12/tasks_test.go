package day12

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 21, Type: _testing.TestData},
			{Path: "day12", Expected: 7653, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 525152, Type: _testing.TestData},
			{Path: "day12", Expected: 60681419004564, Type: _testing.RealData},
		},
	},
}

func TestDay12(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay12(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
