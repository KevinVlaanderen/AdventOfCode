package day5

import (
	"aoc/framework/tasks/tests"
	"go/types"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[string, types.Nil]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[string, types.Nil]{
			{Data: tests.MockData("data"), Expected: "CMZ"},
			{Data: tests.RealData("day5"), Expected: "SHQWSRBDL"},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[string, types.Nil]{
			{Data: tests.MockData("data"), Expected: "MCD"},
			{Data: tests.RealData("day5"), Expected: "CDTQZHBRS"},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
