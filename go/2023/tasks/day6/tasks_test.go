package day6

import (
	"aoc/framework/tasks/tests"
	"go/types"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int, types.Nil]{
			{Data: tests.MockData("data"), Expected: 288},
			{Data: tests.RealData("day6"), Expected: 252000},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int, types.Nil]{
			{Data: tests.MockData("data"), Expected: 71503},
			{Data: tests.RealData("day6"), Expected: 36992486},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
