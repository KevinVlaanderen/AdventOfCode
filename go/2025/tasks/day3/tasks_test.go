package day3

import (
	_testing "aoc/framework/tasks/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, int]{
	{
		Task: Task,
		Tests: []_testing.TestDefinition[int, int]{
			{Path: "data", Param: 2, Expected: 357, Type: _testing.TestData},
			{Path: "day3", Param: 2, Expected: 17109, Type: _testing.RealData},
		},
	},
	{
		Task: Task,
		Tests: []_testing.TestDefinition[int, int]{
			{Path: "data", Param: 12, Expected: 3121910778619, Type: _testing.TestData},
			{Path: "day3", Param: 12, Expected: 169347417057382, Type: _testing.RealData},
		},
	},
}

func Test(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
