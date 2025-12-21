package day21

import (
	"aoc/framework/tasks/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int, int]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int, int]{
			{Path: "data", Expected: 16, Type: tests.TestData, Param: 6},
			{Path: "day21", Expected: 3646, Type: tests.RealData, Param: 64},
		},
	},
	{
		Task:  Task2,
		Tests: []tests.TestDefinition[int, int]{
			// {Path: "data", Expected: 16, Type: tests.TestData, Param: 6},
			// {Path: "data", Expected: 50, Type: tests.TestData, Param: 10},
			// {Path: "data", Expected: 1594, Type: tests.TestData, Param: 50},
			// {Path: "data", Expected: 6536, Type: tests.TestData, Param: 100},
			// {Path: "data", Expected: 167004, Type: tests.TestData, Param: 500},
			// {Path: "data", Expected: 668697, Type: tests.TestData, Param: 1000},
			// {Path: "day21", Expected: -1, Type: tests.RealData, Param: 26501365},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
