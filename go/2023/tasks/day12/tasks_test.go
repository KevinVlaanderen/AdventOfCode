package day12

import (
	"aoc/framework/tasks/tests"
	"go/types"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int, types.Nil]{
			{Data: tests.MockData("data"), Expected: 21},
			{Data: tests.RealData("day12"), Expected: 7653},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int, types.Nil]{
			{Data: tests.MockData("data"), Expected: 525152},
			{Data: tests.RealData("day12"), Expected: 60681419004564},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
