package day10

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int]{
			{"data", 4, _testing.TestData},
			{"data2", 8, _testing.TestData},
			{"day10", 6757, _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int]{
			{"data3", 4, _testing.TestData},
			{"data4", 8, _testing.TestData},
			{"data5", 10, _testing.TestData},
			{"day10", 523, _testing.RealData},
		},
	},
}

func TestDay10(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay10(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
