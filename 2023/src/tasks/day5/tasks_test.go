package day5

import (
	"2023/src/framework/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int]{
			{"data", 35, tests.TestData},
			{"day5", 484023871, tests.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int]{
			{"data", 46, tests.TestData},
			{"day5", 46294175, tests.RealData},
		},
	},
}

func TestDay5(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func BenchmarkDay5(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
