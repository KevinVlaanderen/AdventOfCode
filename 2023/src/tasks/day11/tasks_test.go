package day11

import (
	"2023/src/test"
	"testing"
)

var taskDefinitions = []test.TaskDefinition[int]{
	{
		Task:     Task1,
		TestData: []test.TaskTestDefinition[int]{{"data", 374}},
		RealData: []test.TaskTestDefinition[int]{{"day11", 9418609}},
	},
	{
		Task:     Task2,
		RealData: []test.TaskTestDefinition[int]{{"day11", 593821230983}},
	},
}

func TestDay11(t *testing.T) {
	test.RunTests(t, taskDefinitions)
}

func BenchmarkDay11(b *testing.B) {
	test.RunBenchmarks(b, taskDefinitions)
}
