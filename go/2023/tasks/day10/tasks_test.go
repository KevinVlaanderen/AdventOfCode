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
			{Path: "data", Expected: 4, Type: tests.TestData},
			{Path: "data2", Expected: 8, Type: tests.TestData},
			{Path: "day10", Expected: 6757, Type: tests.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int, types.Nil]{
			{Path: "data3", Expected: 4, Type: tests.TestData},
			{Path: "data4", Expected: 8, Type: tests.TestData},
			{Path: "data5", Expected: 10, Type: tests.TestData},
			{Path: "day10", Expected: 523, Type: tests.RealData},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
