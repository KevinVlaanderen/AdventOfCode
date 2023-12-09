package day1

import (
	"2023/src/test"
	"testing"
)

var taskDefinitions = []test.TaskDefinition[int]{
	{
		Task:     Task1,
		TestData: []test.TaskTestDefinition[int]{{"data", 142}},
		RealData: []test.TaskTestDefinition[int]{{"day1", 54081}},
	},
	{
		Task:     Task2,
		TestData: []test.TaskTestDefinition[int]{{"data2", 281}},
		RealData: []test.TaskTestDefinition[int]{{"day1", 54649}},
	},
}

func TestDay1(t *testing.T) {
	test.RunTests(t, taskDefinitions)
}

func BenchmarkDay1(b *testing.B) {
	test.RunBenchmarks(b, taskDefinitions)
}
