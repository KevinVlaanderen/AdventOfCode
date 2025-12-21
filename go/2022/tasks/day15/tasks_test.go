package day15

import (
	"aoc/framework/tasks/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int, int]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int, int]{
			{Path: "data", Param: 10, Expected: 26, Type: tests.TestData},
			{Path: "day15", Param: 2000000, Expected: 4886370, Type: tests.RealData},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
