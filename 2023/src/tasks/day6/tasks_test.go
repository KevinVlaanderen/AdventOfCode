package day6

import (
	"2023/src/framework/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int]{
			{"data", 288, tests.TestData},
			{"day6", 252000, tests.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int]{
			{"data", 71503, tests.TestData},
			{"day6", 36992486, tests.RealData},
		},
	},
}

func TestDay6(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func BenchmarkDay6(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
