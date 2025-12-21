package day1

import (
	"aoc/framework/tasks/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int, bool]{
	{
		Task: Task,
		Tests: []tests.TestDefinition[int, bool]{
			{Path: "data", Expected: 3, Param: false, Type: tests.TestData},
			{Path: "day1", Expected: 1023, Param: false, Type: tests.RealData},
		},
	},
	{
		Task: Task,
		Tests: []tests.TestDefinition[int, bool]{
			{Path: "data", Expected: 6, Param: true, Type: tests.TestData},
			{Path: "day1", Expected: 5899, Param: true, Type: tests.RealData},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
