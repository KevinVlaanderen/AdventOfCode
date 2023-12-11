package day3

import (
	"2023/src/framework/test"
	"testing"
)

var taskDefinitions = []test.TaskDefinition[int]{
	{
		Task:     Task1,
		TestData: []test.TaskTestDefinition[int]{{"data", 4361}},
		RealData: []test.TaskTestDefinition[int]{{"day3", 546312}},
	},
	{
		Task:     Task2,
		TestData: []test.TaskTestDefinition[int]{{"data", 467835}},
		RealData: []test.TaskTestDefinition[int]{{"day3", 87449461}},
	},
}

func TestDay3(t *testing.T) {
	test.RunTests(t, taskDefinitions)
}

func BenchmarkDay3(b *testing.B) {
	test.RunBenchmarks(b, taskDefinitions)
}
