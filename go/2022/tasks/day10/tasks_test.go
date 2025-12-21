package day10

import (
	"aoc/framework/tasks/tests"
	"go/types"
	"testing"
)

var taskDefinitions1 = []tests.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 13140, Type: tests.TestData},
			{Path: "day10", Expected: 12740, Type: tests.RealData},
		},
	},
}

var taskDefinitions2 = []tests.TaskDefinition[string, types.Nil]{
	{
		Task: Task2,
		Tests: []tests.TestDefinition[string, types.Nil]{
			{Path: "day10", Expected: "RBPARAGF", Type: tests.RealData},
		},
	},
}

func Test1(t *testing.T) {
	tests.RunTests(t, taskDefinitions1)
}

func Test2(t *testing.T) {
	tests.RunTests(t, taskDefinitions2)
}

func Benchmark1(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions1)
}

func Benchmark2(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions2)
}
