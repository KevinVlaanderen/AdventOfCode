package day1

import (
	_testing "aoc/framework/tasks/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 24000, Type: _testing.TestData},
			{Path: "day1", Expected: 67027, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 45000, Type: _testing.TestData},
			{Path: "day1", Expected: 197291, Type: _testing.RealData},
		},
	},
}

func Test(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
