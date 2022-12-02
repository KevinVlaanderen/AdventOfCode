package tasks

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestDay1(t *testing.T) {
	day := Day1{}

	var testDataPath, realDataPath string
	var err error

	if testDataPath, err = createTestDataPath("day1"); err != nil {
		t.Fail()
	}
	if realDataPath, err = createRealDataPath("day1"); err != nil {
		t.Fail()
	}

	t.Run("task 1", func(t *testing.T) {
		t.Run("test data", createTest(day.Task1, testDataPath, 24000))
		t.Run("real data", createTest(day.Task1, realDataPath, 67027))
	})

	t.Run("task 2", func(t *testing.T) {
		t.Run("test data", createTest(day.Task2, testDataPath, 45000))
		t.Run("real data", createTest(day.Task2, realDataPath, 197291))
	})
}

func createTest(task func(string) (*int, error), dataPath string, expected int) func(*testing.T) {
	return func(t *testing.T) {
		var result *int
		var err error

		if result, err = task(dataPath); err != nil {
			t.Fatal(err)
		}

		assertEqual(t, *result, expected)
	}
}

func createTestDataPath(name string) (path string, err error) {
	path, err = filepath.Abs(filepath.Join("testdata", name))
	return
}

func createRealDataPath(name string) (path string, err error) {
	path, err = filepath.Abs(filepath.Join("../../data", name))
	return
}

func loadFile(path string) string {
	content, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Fatal(fmt.Sprintf("Got %v, expected %v", a, b))
}
