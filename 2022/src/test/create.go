//go:build testing

package test

import (
	"2022/src/framework"
	"path/filepath"
	"testing"
)

func CreateTest(task framework.Task, dataPath string, expected int) func(*testing.T) {
	return func(t *testing.T) {
		var result framework.TaskResult

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
