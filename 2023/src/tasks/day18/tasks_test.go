package day18

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int64]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int64]{
			{"data", 62, _testing.TestData},
			{"day18", 70026, _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int64]{
			{"data", 952408144115, _testing.TestData},
			{"day18", 68548301037382, _testing.RealData},
		},
	},
}

func TestDay18(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay18(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
