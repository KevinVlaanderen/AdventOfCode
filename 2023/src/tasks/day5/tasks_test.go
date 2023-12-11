package day5

import (
	"2023/src/framework/test"
	"testing"
)

var taskDefinitions = []test.TaskDefinition[int]{
	{
		Task:     Task1,
		TestData: []test.TaskTestDefinition[int]{{"data", 35}},
		RealData: []test.TaskTestDefinition[int]{{"day5", 484023871}},
	},
	{
		Task:     Task2,
		TestData: []test.TaskTestDefinition[int]{{"data", 46}},
		RealData: []test.TaskTestDefinition[int]{{"day5", 46294175}},
	},
}

func TestDay5(t *testing.T) {
	test.RunTests(t, taskDefinitions)
}

func BenchmarkDay5(b *testing.B) {
	test.RunBenchmarks(b, taskDefinitions)
}
