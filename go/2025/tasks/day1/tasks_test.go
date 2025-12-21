package day1

import (
	"aoc/framework/tasks/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int, bool]{
	{
		Task: Task,
		Tests: []tests.TestDefinition[int, bool]{
			{Data: tests.MockData("data"), Expected: 3, Param: false},
			{Data: tests.RealData("day1"), Expected: 1023, Param: false},
		},
	},
	{
		Task: Task,
		Tests: []tests.TestDefinition[int, bool]{
			{Data: tests.MockData("data"), Expected: 6, Param: true},
			{Data: tests.RealData("day1"), Expected: 5899, Param: true},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
