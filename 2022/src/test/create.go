//go:build testing

package test

import (
	"path/filepath"
	"testing"
)

func CreateTest(task func(string) (*int, error), dataPath string, expected int) func(*testing.T) {
	return func(t *testing.T) {
		var result *int
		var err error

		if result, err = task(dataPath); err != nil {
			t.Fatal(err)
		}

		AssertEqual(t, *result, expected)
	}
}

func CreateTestDataPath(name string) (path string, err error) {
	path, err = filepath.Abs(filepath.Join("testdata", name))
	return
}

func CreateRealDataPath(name string) (path string, err error) {
	path, err = filepath.Abs(filepath.Join("../../data", name))
	return
}
