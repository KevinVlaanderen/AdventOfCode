package day3

import (
	"aoc/framework/tasks/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int, int]{
	{
		Task: Task,
		Tests: []tests.TestDefinition[int, int]{
			{Data: tests.MockData("data"), Param: 2, Expected: 357},
			{Data: tests.RealData("day3"), Param: 2, Expected: 17109},
		},
	},
	{
		Task: Task,
		Tests: []tests.TestDefinition[int, int]{
			{Data: tests.MockData("data"), Param: 12, Expected: 3121910778619},
			{Data: tests.RealData("day3"), Param: 12, Expected: 169347417057382},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
