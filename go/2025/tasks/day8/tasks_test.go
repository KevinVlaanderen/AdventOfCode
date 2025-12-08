package day8

import (
	_testing "aoc/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, int]{
			{Path: "data", Param: 10, Expected: 40, Type: _testing.TestData},
			{Path: "day8", Param: 1000, Expected: 122430, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, int]{
			{Path: "data", Param: 10, Expected: 25272, Type: _testing.TestData},
			{Path: "day8", Param: 1000, Expected: 8135565324, Type: _testing.RealData},
		},
	},
}

func Test(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
