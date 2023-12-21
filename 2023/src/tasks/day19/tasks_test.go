package day19

import (
	_testing "2023/src/framework/testing"
	"go/types"
	"testing"
)

var taskDefinitions = []_testing.TaskDefinition[int, types.Nil]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 19114, Type: _testing.TestData},
			{Path: "day19", Expected: 323625, Type: _testing.RealData},
		},
	},
	{
		Task: Task2,
		Tests: []_testing.TestDefinition[int, types.Nil]{
			{Path: "data", Expected: 167409079868000, Type: _testing.TestData},
			{Path: "day19", Expected: 127447746739409, Type: _testing.RealData},
		},
	},
}

func TestDay19(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func BenchmarkDay19(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
