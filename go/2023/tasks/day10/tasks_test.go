package day10

import (
	"aoc/framework/tasks/tests"
	"go/types"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int, types.Nil]{
			{Data: tests.MockData("data"), Expected: 4},
			{Data: tests.MockData("data2"), Expected: 8},
			{Data: tests.RealData("day10"), Expected: 6757},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int, types.Nil]{
			{Data: tests.MockData("data3"), Expected: 4},
			{Data: tests.MockData("data4"), Expected: 8},
			{Data: tests.MockData("data5"), Expected: 10},
			{Data: tests.RealData("day10"), Expected: 523},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
