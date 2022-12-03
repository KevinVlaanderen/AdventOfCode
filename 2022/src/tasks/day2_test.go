package tasks

import (
	"2022/src/test"
	"testing"
)

func TestDay2(t *testing.T) {
	day := Day2{}

	var testDataPath, realDataPath string
	var err error

	if testDataPath, err = test.CreateTestDataPath("day2"); err != nil {
		t.Fatal(err)
	}
	if realDataPath, err = test.CreateRealDataPath("day2"); err != nil {
		t.Fatal(err)
	}

	t.Run("task 1", func(t *testing.T) {
		t.Run("test data", test.CreateTest(day.Task1, testDataPath, 15))
		t.Run("real data", test.CreateTest(day.Task1, realDataPath, 11475))
	})

	t.Run("task 2", func(t *testing.T) {
		t.Run("test data", test.CreateTest(day.Task2, testDataPath, -1))
		t.Run("real data", test.CreateTest(day.Task2, realDataPath, -1))
	})
}
