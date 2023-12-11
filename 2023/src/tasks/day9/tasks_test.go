package day9

import (
	"2023/src/framework/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int]{
			{"data", 114, tests.TestData},
			{"day9", 1993300041, tests.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int]{
			{"data", 2, tests.TestData},
			{"day9", 1038, tests.RealData},
		},
	},
}

func TestDay9(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func BenchmarkDay9(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
