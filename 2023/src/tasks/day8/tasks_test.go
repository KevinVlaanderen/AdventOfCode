package day8

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int]{
			{"data", 2, _testing.TestData},
			{"data2", 6, _testing.TestData},
			{"day8", 14893, _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int]{
			{"data3", 6, _testing.TestData},
			{"day8", 10241191004509, _testing.RealData},
		},
	},
}

func TestDay8(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay8(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
