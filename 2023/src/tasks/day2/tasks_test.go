package day2

import (
	"2023/src/framework/test"
	"testing"
)

var taskDefinitions = []test.TaskDefinition[int]{
	{
		Task:     Task1,
		TestData: []test.TaskTestDefinition[int]{{"data", 8}},
		RealData: []test.TaskTestDefinition[int]{{"day2", 2278}},
	},
	{
		Task:     Task2,
		TestData: []test.TaskTestDefinition[int]{{"data", 2286}},
		RealData: []test.TaskTestDefinition[int]{{"day2", 67953}},
	},
}

func TestDay2(t *testing.T) {
	test.RunTests(t, taskDefinitions)
}

func BenchmarkDay2(b *testing.B) {
	test.RunBenchmarks(b, taskDefinitions)
}
