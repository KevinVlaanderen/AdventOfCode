package day1

import (
	"2023/src/framework/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int]{
			{"data", 142, tests.TestData},
			{"day1", 54081, tests.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int]{
			{"data2", 281, tests.TestData},
			{"day1", 54649, tests.RealData},
		},
	},
}

func TestDay1(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func BenchmarkDay1(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
