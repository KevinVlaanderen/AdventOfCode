package day13

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int]{
			{"data", 405, _testing.TestData},
			{"day13", 30802, _testing.RealData},
		},
	},
}

func TestDay13(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay13(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
