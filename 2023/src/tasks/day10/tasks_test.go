package day10

import (
	"2023/src/test"
	"testing"
)

var taskDefinitions = []test.TaskDefinition[int]{
	{
		Task: Task1,
		TestData: []test.TaskTestDefinition[int]{
			{"data", 4},
			{"data2", 8},
		},
		RealData: []test.TaskTestDefinition[int]{{"day10", 6757}},
	},
	{
		Task: Task2,
		TestData: []test.TaskTestDefinition[int]{
			{"data3", 4},
			{"data4", 8},
			{"data5", 10},
		},
		RealData: []test.TaskTestDefinition[int]{{"day10", 523}},
	},
}

func TestDay10(t *testing.T) {
	test.RunTests(t, taskDefinitions)
}

func BenchmarkDay10(b *testing.B) {
	test.RunBenchmarks(b, taskDefinitions)
}
