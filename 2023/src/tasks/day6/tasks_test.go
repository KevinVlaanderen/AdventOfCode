package day6

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int]{
			{"data", 288, _testing.TestData},
			{"day6", 252000, _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int]{
			{"data", 71503, _testing.TestData},
			{"day6", 36992486, _testing.RealData},
		},
	},
}

func TestDay6(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay6(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
