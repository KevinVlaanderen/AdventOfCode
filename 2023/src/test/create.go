//go:build testing

package test

import (
	_task "2023/src/framework/task"
	"path/filepath"
	"testing"
)

func CreateTest[T comparable](task _task.Task[T], dataPath string, expected T, skip bool) func(*testing.T) {
	return func(t *testing.T) {
		if skip {
			t.SkipNow()
		}
		var result _task.Result[T]

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
