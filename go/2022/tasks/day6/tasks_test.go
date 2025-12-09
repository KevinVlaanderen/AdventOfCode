package day6

import (
	_testing "aoc/framework/tasks/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data1", Expected: 7, Type: _testing.TestData},
			{Path: "data2", Expected: 5, Type: _testing.TestData},
			{Path: "data3", Expected: 6, Type: _testing.TestData},
			{Path: "data4", Expected: 10, Type: _testing.TestData},
			{Path: "data5", Expected: 11, Type: _testing.TestData},
			{Path: "day6", Expected: 1042, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data1", Expected: 19, Type: _testing.TestData},
			{Path: "data2", Expected: 23, Type: _testing.TestData},
			{Path: "data3", Expected: 23, Type: _testing.TestData},
			{Path: "data4", Expected: 29, Type: _testing.TestData},
			{Path: "data5", Expected: 26, Type: _testing.TestData},
			{Path: "day6", Expected: 2980, Type: _testing.RealData},
		},
	},
}

func Test(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
