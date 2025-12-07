package day6

import (
	_testing "2025/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 4277556, Type: _testing.TestData},
			{Path: "day6", Expected: 4449991244405, Type: _testing.RealData},
		},
	},
}

func TestDay1(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay1(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
