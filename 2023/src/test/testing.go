//go:build testing

package test

import (
	"2023/src/framework"
	"fmt"
	"path/filepath"
	"testing"
)

type TaskDefinition[T comparable] struct {
	Task     framework.Task[T]
	TestData []TaskTestDefinition[T]
	RealData []TaskTestDefinition[T]
}

type TaskTestDefinition[T comparable] struct {
	Path     string
	Expected T
}

func RunTests(t *testing.T, taskDefinitions []TaskDefinition[int]) {
	for taskIndex, taskDefinition := range taskDefinitions {
		t.Run(fmt.Sprintf("Task%v", taskIndex+1), func(tTask *testing.T) {
			for _, testDefinition := range taskDefinition.TestData {
				if path, err := CreateTestDataPath(testDefinition.Path); err != nil {
					tTask.Fatal(err)
				} else {
					test := CreateTest(taskDefinition.Task, path, testDefinition.Expected)
					tTask.Run(fmt.Sprintf("TestData(%v)", testDefinition.Path), test)
				}
			}
			for _, testDefinition := range taskDefinition.RealData {
				if path, err := CreateRealDataPath(testDefinition.Path); err != nil {
					tTask.Fatal(err)
				} else {
					test := CreateTest(taskDefinition.Task, path, testDefinition.Expected)
					tTask.Run(fmt.Sprintf("RealData(%v)", testDefinition.Path), test)
				}
			}
		})
	}
}

func RunBenchmarks(b *testing.B, taskDefinitions []TaskDefinition[int]) {
	for taskIndex, taskDefinition := range taskDefinitions {
		b.Run(fmt.Sprintf("Task%v", taskIndex+1), func(bTask *testing.B) {
			for _, testDefinition := range taskDefinition.TestData {
				if path, err := CreateTestDataPath(testDefinition.Path); err != nil {
					bTask.Fatal(err)
				} else {
					test := CreateBenchmark(taskDefinition.Task, path)
					bTask.Run(fmt.Sprintf("TestData(%v)", testDefinition.Path), test)
				}
			}
			for _, testDefinition := range taskDefinition.RealData {
				if path, err := CreateRealDataPath(testDefinition.Path); err != nil {
					bTask.Fatal(err)
				} else {
					test := CreateBenchmark(taskDefinition.Task, path)
					bTask.Run(fmt.Sprintf("RealData(%v)", testDefinition.Path), test)
				}
			}
		})
	}
}

func CreateTest[T comparable](task framework.Task[T], dataPath string, expected T) func(*testing.T) {
	return func(t *testing.T) {
		if result := task(dataPath); result.Error != nil {
			t.Fatal(result.Error)
		} else {
			AssertEqual(t, result.Value, expected)
		}
	}
}

func CreateBenchmark[T comparable](task framework.Task[T], dataPath string) func(*testing.B) {
	return func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			task(dataPath)
		}
	}
}

func CreateTestDataPath(name string) (path string, err error) {
	path, err = filepath.Abs(filepath.Join("testdata", name))
	return
}

func CreateRealDataPath(name string) (path string, err error) {
	path, err = filepath.Abs(filepath.Join("../../../data", name))
	return
}
