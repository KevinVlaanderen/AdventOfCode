package day17

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task:  Task1,
		Tests: []_testing.TestDefinition[int]{
			//{"data", 102, _testing.TestData},
			//{"day17", -1, _testing.RealData},
		},
	},
}

func TestDay17(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay17(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
