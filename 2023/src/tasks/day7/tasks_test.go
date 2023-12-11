package day7

import (
	"2023/src/framework/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int]{
			{"data", 6440, tests.TestData},
			{"day7", 251216224, tests.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int]{
			{"data", 5905, tests.TestData},
			{"day7", 250825971, tests.RealData},
		},
	},
}

func TestDay7(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func BenchmarkDay7(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
