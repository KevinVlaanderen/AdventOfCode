package day7

import (
	"2023/src/test"
	"testing"
)

var taskDefinitions = []test.TaskDefinition[int]{
	{
		Task:     Task1,
		TestData: []test.TaskTestDefinition[int]{{"data", 6440}},
		RealData: []test.TaskTestDefinition[int]{{"day7", 251216224}},
	},
	{
		Task:     Task2,
		TestData: []test.TaskTestDefinition[int]{{"data", 5905}},
		RealData: []test.TaskTestDefinition[int]{{"day7", 250825971}},
	},
}

func TestDay7(t *testing.T) {
	test.RunTests(t, taskDefinitions)
}

func BenchmarkDay7(b *testing.B) {
	test.RunBenchmarks(b, taskDefinitions)
}
