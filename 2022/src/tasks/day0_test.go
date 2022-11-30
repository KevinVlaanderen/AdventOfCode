package tasks

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestDay0(t *testing.T) {
	t.Run("task 1", func(t *testing.T) {
		t.Run("test data", createTest(Day0{}.task1, testDataPath("day0"), 0))
		t.Run("real data", createTest(Day0{}.task1, realDataPath("day0"), 1))
	})
	t.Run("task 2", func(t *testing.T) {
		t.Run("test data", createTest(Day0{}.task1, testDataPath("day0"), 0))
		t.Run("real data", createTest(Day0{}.task1, realDataPath("day0"), 1))
	})
}

func createTest(task func(string) int, dataPath string, expected int) func(*testing.T) {
	return func(t *testing.T) {
		data := loadFile(dataPath)
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
