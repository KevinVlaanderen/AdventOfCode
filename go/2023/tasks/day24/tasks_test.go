package day24

import (
	_testing "aoc/framework/testing"
	"testing"

	"github.com/samber/lo"
)

var taskDefinitions = []_testing.TaskDefinition[int, lo.Tuple2[int, int]]{
	{
		Task: Task1,
		Tests: []_testing.TestDefinition[int, lo.Tuple2[int, int]]{
			{Path: "data", Expected: 2, Type: _testing.TestData, Param: lo.Tuple2[int, int]{A: 7, B: 27}},
			{Path: "day24", Expected: 21843, Type: _testing.RealData, Param: lo.Tuple2[int, int]{A: 200000000000000, B: 400000000000000}},
		},
	},
}

func Test(t *testing.T) {
	_testing.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	_testing.RunBenchmarks(b, taskDefinitions)
}
