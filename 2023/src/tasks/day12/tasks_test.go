package day12

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int]{
			{"data", 21, _testing.TestData},
			{"day12", 7653, _testing.RealData},
		},
	},
}

func TestDay12(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay12(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
