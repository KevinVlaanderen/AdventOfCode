package day20

import (
	"aoc/framework/tasks/tests"
	"go/types"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 32000000, Type: tests.TestData},
			{Path: "data2", Expected: 11687500, Type: tests.TestData},
			{Path: "day20", Expected: 821985143, Type: tests.RealData},
		},
	},
	{
		Task:  Task2,
		Tests: []tests.TestDefinition[int, types.Nil]{
			// {Path: "day20", Expected: -1, Type: tests.RealData},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
