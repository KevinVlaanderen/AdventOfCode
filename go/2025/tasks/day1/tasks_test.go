package day1

import (
	_testing "aoc/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, bool]{
	{
		Task: Task,
		Tests: []_testing.TestDefinition[int, bool]{
			{Path: "data", Expected: 3, Param: false, Type: _testing.TestData},
			{Path: "day1", Expected: 1023, Param: false, Type: _testing.RealData},
		},
	},
	{
		Task: Task,
		Tests: []_testing.TestDefinition[int, bool]{
			{Path: "data", Expected: 6, Param: true, Type: _testing.TestData},
			{Path: "day1", Expected: 5899, Param: true, Type: _testing.RealData},
		},
	},
}

func Test(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
