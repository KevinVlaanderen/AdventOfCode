package day8

import (
	_testing "2025/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, int]{
			{Path: "data", Param: 10, Expected: 40, Type: _testing.TestData},
			{Path: "day8", Param: 1000, Expected: 122430, Type: _testing.RealData},
		},
	},
}

func TestDay1(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay1(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
