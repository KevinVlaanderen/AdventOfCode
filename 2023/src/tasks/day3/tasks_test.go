package day3

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int]{
			{"data", 4361, _testing.TestData},
			{"day3", 546312, _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int]{
			{"data", 467835, _testing.TestData},
			{"day3", 87449461, _testing.RealData},
		},
	},
}

func TestDay3(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay3(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
