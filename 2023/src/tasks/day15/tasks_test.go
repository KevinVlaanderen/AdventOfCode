package day15

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int]{
			{"data", 52, _testing.TestData},
			{"data2", 1320, _testing.TestData},
			{"day15", 505427, _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int]{
			{"data2", 145, _testing.TestData},
			{"day15", 243747, _testing.RealData},
		},
	},
}

func TestDay15(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay15(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
