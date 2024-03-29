//go:build testing

package test

import (
	"2022/src/framework/tasks"
	"path/filepath"
	"testing"
)

func CreateTest[T comparable](task tasks.Task[T], dataPath string, expected T) func(*testing.T) {
	return func(t *testing.T) {
		var result tasks.TaskResult[T]

		if result = task(dataPath); result.Error != nil {
			t.Fatal(result.Error)
		}

		AssertEqual(t, result.Value, expected)
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
