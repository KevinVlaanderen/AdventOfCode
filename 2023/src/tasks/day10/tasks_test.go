package day10

import (
	"2023/src/framework/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int]{
			{"data", 4, tests.TestData},
			{"data2", 8, tests.TestData},
			{"day10", 6757, tests.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int]{
			{"data3", 4, tests.TestData},
			{"data4", 8, tests.TestData},
			{"data5", 10, tests.TestData},
			{"day10", 523, tests.RealData},
		},
	},
}

func TestDay10(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func BenchmarkDay10(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
