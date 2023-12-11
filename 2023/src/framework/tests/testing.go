//go:build testing

package tests

import (
	"2023/src/framework"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type TestType int

const (
	TestData TestType = iota
	RealData
)

type TaskDefinition[T comparable] struct {
	Task  framework.Task[T]
	Tests []TestDefinition[T]
}

type TestDefinition[T comparable] struct {
	Path     string
	Expected T
	Type     TestType
}

func (t TestDefinition[T]) name() string {
	switch t.Type {
	case TestData:
		return fmt.Sprintf("TestData(%v)", t.Path)
	case RealData:
		return fmt.Sprintf("RealData(%v)", t.Path)
	}
	panic(fmt.Errorf("unknown test type %v", t.Type))
}

func (t TestDefinition[T]) dataPath() string {
	var path string
	var err error

	switch t.Type {
	case TestData:
		path, err = CreateTestDataPath(t.Path)
	case RealData:
		path, err = CreateRealDataPath(t.Path)
	default:
		panic(fmt.Errorf("unknown test type %v", t.Type))

	}
	if err != nil {
		panic(err)
	} else {
		return path
	}
}

func RunTests[T comparable](t *testing.T, taskDefinitions []TaskDefinition[T]) {
	t.Parallel()

	for taskIndex, taskDefinition := range taskDefinitions {
		t.Run(fmt.Sprintf("Task%v", taskIndex+1), func(tTask *testing.T) {
			for _, testDefinition := range taskDefinition.Tests {
				runTest(tTask, taskDefinition, testDefinition)
			}
		})
	}
}

func runTest[T comparable](t *testing.T, taskDefinition TaskDefinition[T], testDefinition TestDefinition[T]) {
	data := strings.TrimSpace(ReadAll(testDefinition.dataPath()))
	test := CreateTest(taskDefinition.Task, data, testDefinition.Expected)
	t.Run(testDefinition.name(), test)
}

func RunBenchmarks(b *testing.B, taskDefinitions []TaskDefinition[int]) {
	for taskIndex, taskDefinition := range taskDefinitions {
		b.Run(fmt.Sprintf("Task%v", taskIndex+1), func(bTask *testing.B) {
			for _, testDefinition := range taskDefinition.Tests {
				runBenchmark(bTask, taskDefinition, testDefinition)
			}
		})
	}
}

func runBenchmark[T comparable](b *testing.B, taskDefinition TaskDefinition[T], testDefinition TestDefinition[T]) {
	data := strings.TrimSpace(ReadAll(testDefinition.dataPath()))
	test := CreateBenchmark(taskDefinition.Task, data)
	b.Run(testDefinition.name(), test)
}

func CreateTest[T comparable](task framework.Task[T], data string, expected T) func(*testing.T) {
	return func(t *testing.T) {
		if result := task(data); result.Error != nil {
			t.Fatal(result.Error)
		} else {
			AssertEqual(t, result.Value, expected)
		}
	}
}

func CreateBenchmark[T comparable](task framework.Task[T], data string) func(*testing.B) {
	return func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			task(data)
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

func ReadAll(path string) string {
	if bytes, err := os.ReadFile(path); err != nil {
		panic(err)
	} else {
		return string(bytes)
	}
}
