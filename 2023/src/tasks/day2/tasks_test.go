package day2

import (
	"2023/src/framework/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int]{
			{"data", 8, tests.TestData},
			{"day2", 2278, tests.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int]{
			{"data", 2286, tests.TestData},
			{"day2", 67953, tests.RealData},
		},
	},
}

func TestDay2(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func BenchmarkDay2(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
