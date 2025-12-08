package day15

import (
	_testing "aoc/framework/testing"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, int]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, int]{
			{Path: "data", Param: 10, Expected: 26, Type: _testing.TestData},
			{Path: "day15", Param: 2000000, Expected: 4886370, Type: _testing.RealData},
		},
	},
}

func Test(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
