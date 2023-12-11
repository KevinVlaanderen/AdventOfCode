package day7

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int]{
			{"data", 6440, _testing.TestData},
			{"day7", 251216224, _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int]{
			{"data", 5905, _testing.TestData},
			{"day7", 250825971, _testing.RealData},
		},
	},
}

func TestDay7(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay7(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
