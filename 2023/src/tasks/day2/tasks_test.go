package day2

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int]{
			{"data", 8, _testing.TestData},
			{"day2", 2278, _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int]{
			{"data", 2286, _testing.TestData},
			{"day2", 67953, _testing.RealData},
		},
	},
}

func TestDay2(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay2(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
