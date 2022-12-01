package tasks

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestDay1(t *testing.T) {
	t.Run("task 1", func(t *testing.T) {
		t.Run("test data", createTest(Day1{}.task1, testDataPath("day1"), 24000))
		t.Run("real data", createTest(Day1{}.task1, realDataPath("day1"), 67027))
	})
}

func createTest(task func(string) int, dataPath string, expected int) func(*testing.T) {
	return func(t *testing.T) {
		data := loadFile(dataPath)
		result := task(data)

		assertEqual(t, result, expected)
		if task(data) != expected {
			t.Fail()
		}
	}
}

func testDataPath(name string) (path string) {
	path = filepath.Join("testdata", name)
	return
}

func realDataPath(name string) (path string) {
	path = filepath.Join("../../data", name)
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
