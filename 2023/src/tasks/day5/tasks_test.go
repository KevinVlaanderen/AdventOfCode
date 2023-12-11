package day5

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int]{
			{"data", 35, _testing.TestData},
			{"day5", 484023871, _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int]{
			{"data", 46, _testing.TestData},
			{"day5", 46294175, _testing.RealData},
		},
	},
}

func TestDay5(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay5(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
