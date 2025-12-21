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
			{Data: tests.MockData("data"), Expected: 32000000},
			{Data: tests.MockData("data2"), Expected: 11687500},
			{Data: tests.RealData("day20"), Expected: 821985143},
		},
	},
	{
		Task:  Task2,
		Tests: []tests.TestDefinition[int, types.Nil]{
			// {Data: tests.RealData("day20"), Expected: -1},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
