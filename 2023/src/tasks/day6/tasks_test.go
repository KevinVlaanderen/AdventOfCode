package day6

import (
	"2023/src/test"
	"testing"
)

var taskDefinitions = []test.TaskDefinition[int]{
	{
		Task:     Task1,
		TestData: []test.TaskTestDefinition[int]{{"data", 288}},
		RealData: []test.TaskTestDefinition[int]{{"day6", 252000}},
	},
	{
		Task:     Task2,
		TestData: []test.TaskTestDefinition[int]{{"data", 71503}},
		RealData: []test.TaskTestDefinition[int]{{"day6", 36992486}},
	},
}

func TestDay6(t *testing.T) {
	test.RunTests(t, taskDefinitions)
}

func BenchmarkDay6(b *testing.B) {
	test.RunBenchmarks(b, taskDefinitions)
}
