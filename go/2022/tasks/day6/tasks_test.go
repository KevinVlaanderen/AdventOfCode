package day6

import (
	"aoc/framework/tasks/tests"
	"go/types"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int, types.Nil]{
			{Path: "data1", Expected: 7, Type: tests.TestData},
			{Path: "data2", Expected: 5, Type: tests.TestData},
			{Path: "data3", Expected: 6, Type: tests.TestData},
			{Path: "data4", Expected: 10, Type: tests.TestData},
			{Path: "data5", Expected: 11, Type: tests.TestData},
			{Path: "day6", Expected: 1042, Type: tests.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int, types.Nil]{
			{Path: "data1", Expected: 19, Type: tests.TestData},
			{Path: "data2", Expected: 23, Type: tests.TestData},
			{Path: "data3", Expected: 23, Type: tests.TestData},
			{Path: "data4", Expected: 29, Type: tests.TestData},
			{Path: "data5", Expected: 26, Type: tests.TestData},
			{Path: "day6", Expected: 2980, Type: tests.RealData},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
