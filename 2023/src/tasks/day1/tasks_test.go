package day1

import (
	"2023/src/framework/test"
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

var data map[string]string

func ReadData(taskDefinitions []test.TaskDefinition[int]) (data map[string]string) {
	for _, taskDefinition := range taskDefinitions {
		for _, testDataDefinition := range taskDefinition.TestData {
			data[testDataDefinition.Path] = test.ReadAll(testDataDefinition.Path)
		}
		for _, realDataDefinition := range taskDefinition.RealData {
			data[realDataDefinition.Path] = test.ReadAll(realDataDefinition.Path)
		}
	}
	return
}

func TestDay1(t *testing.T) {
	test.RunTests(t, taskDefinitions)
}

func BenchmarkDay1(b *testing.B) {
	test.RunBenchmarks(b, taskDefinitions)
}
