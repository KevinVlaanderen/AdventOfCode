package day15

import (
	"aoc/framework/tasks/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int, int]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int, int]{
			{Data: tests.MockData("data"), Param: 10, Expected: 26},
			{Data: tests.RealData("day15"), Param: 2000000, Expected: 4886370},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
