package day10

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 4, Type: _testing.TestData},
			{Path: "data2", Expected: 8, Type: _testing.TestData},
			{Path: "day10", Expected: 6757, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data3", Expected: 4, Type: _testing.TestData},
			{Path: "data4", Expected: 8, Type: _testing.TestData},
			{Path: "data5", Expected: 10, Type: _testing.TestData},
			{Path: "day10", Expected: 523, Type: _testing.RealData},
		},
	},
}

func TestDay10(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay10(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
