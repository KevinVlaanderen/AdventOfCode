package day24

import (
	"aoc/framework/tasks/tests"
	"testing"

	"github.com/samber/lo"
)

var taskDefinitions = []tests.TaskDefinition[int, lo.Tuple2[int, int]]{
	{
		Task: Task1,
		Tests: []tests.TestDefinition[int, lo.Tuple2[int, int]]{
			{Path: "data", Expected: 2, Type: tests.TestData, Param: lo.Tuple2[int, int]{A: 7, B: 27}},
			{Path: "day24", Expected: 21843, Type: tests.RealData, Param: lo.Tuple2[int, int]{A: 200000000000000, B: 400000000000000}},
		},
	},
}

func Test(t *testing.T) {
	tests.RunTests(t, taskDefinitions)
}

func Benchmark(b *testing.B) {
	tests.RunBenchmarks(b, taskDefinitions)
}
