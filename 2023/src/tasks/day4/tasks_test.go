package day4

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int]{
			{"data", 13, _testing.TestData},
			{"day4", 22193, _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int]{
			{"data", 30, _testing.TestData},
			{"day4", 5625994, _testing.RealData},
		},
	},
}

func TestDay4(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay4(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
