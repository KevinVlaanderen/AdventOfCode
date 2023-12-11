package day8

import (
	"2023/src/framework/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int]{
			{"data", 2, tests.TestData},
			{"data2", 6, tests.TestData},
			{"day8", 14893, tests.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int]{
			{"data3", 6, tests.TestData},
			{"day8", 10241191004509, tests.RealData},
		},
	},
}

func TestDay8(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func BenchmarkDay8(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
