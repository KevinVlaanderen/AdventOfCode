package day1

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int]{
			{"data", 142, _testing.TestData},
			{"day1", 54081, _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int]{
			{"data2", 281, _testing.TestData},
			{"day1", 54649, _testing.RealData},
		},
	},
}

func TestDay1(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay1(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
