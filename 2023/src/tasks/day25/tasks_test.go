package day25

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 54, Type: _testing.TestData},
			{Path: "day25", Expected: -1, Type: _testing.RealData},
		},
	},
}

func TestDay25(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay25(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
