package day19

import (
	"aoc/framework/tasks/tests"
	"go/types"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int, types.Nil]{
			{Data: tests.MockData("data"), Expected: 19114},
			{Data: tests.RealData("day19"), Expected: 323625},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int, types.Nil]{
			{Data: tests.MockData("data"), Expected: 167409079868000},
			{Data: tests.RealData("day19"), Expected: 127447746739409},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
