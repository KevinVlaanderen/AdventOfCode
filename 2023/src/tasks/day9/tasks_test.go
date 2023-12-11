package day9

import (
	"2023/src/framework/test"
	"testing"
)

var taskDefinitions = []test.TaskDefinition[int]{
	{
		Task:     Task1,
		TestData: []test.TaskTestDefinition[int]{{"data", 114}},
		RealData: []test.TaskTestDefinition[int]{{"day9", 1993300041}},
	},
	{
		Task:     Task2,
		TestData: []test.TaskTestDefinition[int]{{"data", 2}},
		RealData: []test.TaskTestDefinition[int]{{"day9", 1038}},
	},
}

func TestDay9(t *testing.T) {
	test.RunTests(t, taskDefinitions)
}

func BenchmarkDay9(b *testing.B) {
	test.RunBenchmarks(b, taskDefinitions)
}
