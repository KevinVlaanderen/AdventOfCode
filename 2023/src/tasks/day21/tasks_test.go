package day21

import (
	_testing "2023/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, int]{
			{Path: "data", Expected: 16, Type: _testing.TestData, Param: 6},
			{Path: "day21", Expected: 3646, Type: _testing.RealData, Param: 64},
		},
	},
}

func TestDay21(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay21(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
