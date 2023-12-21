package day20

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 32000000, Type: _testing.TestData},
			{Path: "data2", Expected: 11687500, Type: _testing.TestData},
			{Path: "day20", Expected: 821985143, Type: _testing.RealData},
		},
	},
	{
		Task:  Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			//{Path: "day20", Expected: -1, Type: _testing.RealData},
		},
	},
}

func TestDay20(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay20(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
