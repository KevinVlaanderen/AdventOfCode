package day4

import (
	"2023/src/framework/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int]{
			{"data", 13, tests.TestData},
			{"day4", 22193, tests.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int]{
			{"data", 30, tests.TestData},
			{"day4", 5625994, tests.RealData},
		},
	},
}

func TestDay4(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func BenchmarkDay4(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
