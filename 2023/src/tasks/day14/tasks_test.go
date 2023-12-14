package day14

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int]{
			{"data", 136, _testing.TestData},
			{"day14", 113525, _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int]{
			{"data", 64, _testing.TestData},
			{"day14", 101292, _testing.RealData},
		},
	},
}

func TestDay14(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay14(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
