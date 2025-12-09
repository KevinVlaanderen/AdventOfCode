package day5

import (
	_testing "aoc/framework/tasks/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[string, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[string, types.Nil]{
			{Path: "data", Expected: "CMZ", Type: _testing.TestData},
			{Path: "day5", Expected: "SHQWSRBDL", Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[string, types.Nil]{
			{Path: "data", Expected: "MCD", Type: _testing.TestData},
			{Path: "day5", Expected: "CDTQZHBRS", Type: _testing.RealData},
		},
	},
}

func Test(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
