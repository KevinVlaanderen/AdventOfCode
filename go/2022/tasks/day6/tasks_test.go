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
			{Data: tests.MockData("data1"), Expected: 7},
			{Data: tests.MockData("data2"), Expected: 5},
			{Data: tests.MockData("data3"), Expected: 6},
			{Data: tests.MockData("data4"), Expected: 10},
			{Data: tests.MockData("data5"), Expected: 11},
			{Data: tests.RealData("day6"), Expected: 1042},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int, types.Nil]{
			{Data: tests.MockData("data1"), Expected: 19},
			{Data: tests.MockData("data2"), Expected: 23},
			{Data: tests.MockData("data3"), Expected: 23},
			{Data: tests.MockData("data4"), Expected: 29},
			{Data: tests.MockData("data5"), Expected: 26},
			{Data: tests.RealData("day6"), Expected: 2980},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
