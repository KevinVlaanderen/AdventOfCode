package day8

import (
	"2023/src/test"
	"testing"
)

var taskDefinitions = []test.TaskDefinition[int]{
	{
		Task: Task1,
		TestData: []test.TaskTestDefinition[int]{
			{"data", 2},
			{"data2", 6},
		},
		RealData: []test.TaskTestDefinition[int]{{"day8", 14893}},
	},
	{
		Task:     Task2,
		TestData: []test.TaskTestDefinition[int]{{"data3", 6}},
		RealData: []test.TaskTestDefinition[int]{{"day8", 10241191004509}},
	},
}

func TestDay8(t *testing.T) {
	test.RunTests(t, taskDefinitions)
}

func BenchmarkDay8(b *testing.B) {
	test.RunBenchmarks(b, taskDefinitions)
}
