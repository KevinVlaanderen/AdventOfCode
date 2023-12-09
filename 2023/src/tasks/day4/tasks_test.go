package day4

import (
	"2023/src/test"
	"testing"
)

var taskDefinitions = []test.TaskDefinition[int]{
	{
		Task:     Task1,
		TestData: []test.TaskTestDefinition[int]{{"data", 13}},
		RealData: []test.TaskTestDefinition[int]{{"day4", 22193}},
	},
	{
		Task:     Task2,
		TestData: []test.TaskTestDefinition[int]{{"data", 30}},
		RealData: []test.TaskTestDefinition[int]{{"day4", 5625994}},
	},
}

func TestDay4(t *testing.T) {
	test.RunTests(t, taskDefinitions)
}

func BenchmarkDay4(b *testing.B) {
	test.RunBenchmarks(b, taskDefinitions)
}
