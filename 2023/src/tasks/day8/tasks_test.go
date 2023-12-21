package day8

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 2, Type: _testing.TestData},
			{Path: "data2", Expected: 6, Type: _testing.TestData},
			{Path: "day8", Expected: 14893, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data3", Expected: 6, Type: _testing.TestData},
			{Path: "day8", Expected: 10241191004509, Type: _testing.RealData},
		},
	},
}

func TestDay8(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay8(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
