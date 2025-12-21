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
			{Path: "data", Expected: "CMZ", Type: tests.TestData},
			{Path: "day5", Expected: "SHQWSRBDL", Type: tests.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[string, types.Nil]{
			{Path: "data", Expected: "MCD", Type: tests.TestData},
			{Path: "day5", Expected: "CDTQZHBRS", Type: tests.RealData},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
