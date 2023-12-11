package day11

import (
	"2023/src/framework/tests"
	"testing"
)

var taskDefinitions = []tests.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int]{
			{"data", 374, tests.TestData},
			{"day11", 9418609, tests.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []tests.TestDefinition[int]{
			{"day11", 593821230983, tests.RealData},
		},
	},
}

func TestDay11(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func BenchmarkDay11(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
