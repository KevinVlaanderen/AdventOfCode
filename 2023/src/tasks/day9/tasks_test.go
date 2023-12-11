package day9

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int]{
			{"data", 114, _testing.TestData},
			{"day9", 1993300041, _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int]{
			{"data", 2, _testing.TestData},
			{"day9", 1038, _testing.RealData},
		},
	},
}

func TestDay9(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay9(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
