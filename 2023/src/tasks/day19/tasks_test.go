package day19

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int]{
			{"data", 19114, _testing.TestData},
			{"day19", 323625, _testing.RealData},
		},
	},
}

func TestDay19(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay19(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
