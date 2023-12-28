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
	{
		Task:  Task2,
		Tests: []_testing.TestDefinition[int, int]{
			//{Path: "data", Expected: 16, Type: _testing.TestData, Param: 6},
			//{Path: "data", Expected: 50, Type: _testing.TestData, Param: 10},
			//{Path: "data", Expected: 1594, Type: _testing.TestData, Param: 50},
			//{Path: "data", Expected: 6536, Type: _testing.TestData, Param: 100},
			//{Path: "data", Expected: 167004, Type: _testing.TestData, Param: 500},
			//{Path: "data", Expected: 668697, Type: _testing.TestData, Param: 1000},
			//{Path: "day21", Expected: -1, Type: _testing.RealData, Param: 26501365},
		},
	},
}

func TestDay21(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay21(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
