package day1

import (
	_testing "2025/src/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, bool]{
	{
		Task: Task,
		Tests: []_testing.TestDefinition[int, bool]{
			{Path: "data", Expected: 3, Param: false, Type: _testing.TestData},
			{Path: "day1", Expected: 1023, Param: false, Type: _testing.RealData},
		},
	},
	{
		Task: Task,
		Tests: []_testing.TestDefinition[int, bool]{
			{Path: "data", Expected: 6, Param: true, Type: _testing.TestData},
			{Path: "day1", Expected: 5899, Param: true, Type: _testing.RealData},
		},
	},
}

func TestDay1(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay1(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
