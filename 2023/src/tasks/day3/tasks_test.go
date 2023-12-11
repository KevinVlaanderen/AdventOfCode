package day3

import (
	"2023/src/framework/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int]{
			{"data", 4361, tests.TestData},
			{"day3", 546312, tests.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int]{
			{"data", 467835, tests.TestData},
			{"day3", 87449461, tests.RealData},
		},
	},
}

func TestDay3(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func BenchmarkDay3(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
