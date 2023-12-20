package day20

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int]{
			{"data", 32000000, _testing.TestData},
			{"data2", 11687500, _testing.TestData},
			{"day20", 821985143, _testing.RealData},
		},
	},
	{
		Task:  Task2,
		Tests: []_testing.TestDefinition[int]{
			//{"day20", -1, _testing.RealData},
		},
	},
}

func TestDay20(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay20(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
