package day9

import (
	"aoc/framework/tasks/tests"
	"go/types"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int, types.Nil]{
			{Data: tests.MockData("data1"), Expected: 13},
			{Data: tests.RealData("day9"), Expected: 6470},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int, types.Nil]{
			{Data: tests.MockData("data1"), Expected: 1},
			{Data: tests.MockData("data2"), Expected: 36},
			{Data: tests.RealData("day9"), Expected: 2658},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
