package day8

import (
	"aoc/framework/tasks/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int, int]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int, int]{
			{Data: tests.MockData("data"), Param: 10, Expected: 40},
			{Data: tests.RealData("day8"), Param: 1000, Expected: 122430},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int, int]{
			{Data: tests.MockData("data"), Param: 10, Expected: 25272},
			{Data: tests.RealData("day8"), Param: 1000, Expected: 8135565324},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
