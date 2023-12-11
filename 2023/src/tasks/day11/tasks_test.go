package day11

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int]{
			{"data", 374, _testing.TestData},
			{"day11", 9418609, _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int]{
			{"day11", 593821230983, _testing.RealData},
		},
	},
}

func TestDay11(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay11(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
