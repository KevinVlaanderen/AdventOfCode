package tests

import (
	"aoc/framework/tasks"
	"fmt"
	"testing"
)

type TaskDefinition[T comparable, P any] struct {
	Task  tasks.Task[T, P]
	Tests []TestDefinition[T, P]
}

type TestDefinition[T comparable, P any] struct {
	Data     DataDescriptor
	Expected T
	Param    P
}

func (t TestDefinition[T, P]) name() string {
	switch t.Data.dataType {
	case Mock:
		return fmt.Sprintf("MockData(%v)", t.Data.path)
	case Real:
		return fmt.Sprintf("RealData(%v)", t.Data.path)
	}
	panic(fmt.Errorf("unknown test type %v", t.Data.dataType))
}

func RunTests[T comparable, P any](t *testing.T, taskDefinitions []TaskDefinition[T, P]) {
	t.Parallel()

	for taskIndex, taskDefinition := range taskDefinitions {
		t.Run(fmt.Sprintf("Task%v", taskIndex+1), func(tTask *testing.T) {
			for _, testDefinition := range taskDefinition.Tests {
				if err := runTest(tTask, taskDefinition, testDefinition); err != nil {
					tTask.Error(err)
				}
			}
		})
	}
}

func runTest[T comparable, P any](t *testing.T, taskDefinition TaskDefinition[T, P], testDefinition TestDefinition[T, P]) error {
	data, err := ReadData(testDefinition.Data)
	if err != nil {
		return err
	}

	test := CreateTest(taskDefinition.Task, data, testDefinition.Param, testDefinition.Expected)
	t.Run(testDefinition.name(), test)

	return nil
}

func RunBenchmarks[T comparable, P any](b *testing.B, taskDefinitions []TaskDefinition[T, P]) {
	for taskIndex, taskDefinition := range taskDefinitions {
		b.Run(fmt.Sprintf("Task%v", taskIndex+1), func(bTask *testing.B) {
			for _, testDefinition := range taskDefinition.Tests {
				if err := runBenchmark(bTask, taskDefinition, testDefinition); err != nil {
					bTask.Error(err)
				}
			}
		})
	}
}

func runBenchmark[T comparable, P any](b *testing.B, taskDefinition TaskDefinition[T, P], testDefinition TestDefinition[T, P]) error {
	data, err := ReadData(testDefinition.Data)
	if err != nil {
		return err
	}

	test := CreateBenchmark(taskDefinition.Task, data, testDefinition.Param)
	b.Run(testDefinition.name(), test)

	return nil
}

func CreateTest[T comparable, P any](task tasks.Task[T, P], data string, param P, expected T) func(*testing.T) {
	return func(t *testing.T) {
		if result := task(data, param); result.Error != nil {
			t.Fatal(result.Error)
		} else {
			AssertEqual(t, result.Value, expected)
		}
	}
}

func CreateBenchmark[T comparable, P any](task tasks.Task[T, P], data string, param P) func(*testing.B) {
	return func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			task(data, param)
		}
	}
}
