package day16

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int]{
			{"data", 46, _testing.TestData},
			{"day16", 7210, _testing.RealData},
		},
	},
}

func TestDay16(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay16(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
