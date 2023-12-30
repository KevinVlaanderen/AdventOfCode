package day23

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 94, Type: _testing.TestData},
			{Path: "day23", Expected: 2034, Type: _testing.RealData},
		},
	},
}

func TestDay23(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay23(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
