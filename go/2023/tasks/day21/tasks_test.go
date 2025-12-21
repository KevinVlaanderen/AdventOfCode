package day21

import (
	"aoc/framework/tasks/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int, int]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int, int]{
			{Data: tests.MockData("data"), Param: 6, Expected: 16},
			{Data: tests.RealData("day21"), Param: 64, Expected: 3646},
		},
	},
	{
		Task:  Task2,
		Tests: []tests.TestDefinition[int, int]{
			// {Data: tests.MockData("data"), Expected: 16},
			// {Data: tests.MockData("data"), Expected: 50},
			// {Data: tests.MockData("data"), Expected: 1594},
			// {Data: tests.MockData("data"), Expected: 6536},
			// {Data: tests.MockData("data"), Expected: 167004},
			// {Data: tests.MockData("data"), Expected: 668697},
			// {Data: tests.RealData("day21"), Expected: -1},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
