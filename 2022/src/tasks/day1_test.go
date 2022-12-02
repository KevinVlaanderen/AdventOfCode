package tasks

import (
	"2022/src/test"
	"testing"
)

func TestDay1(t *testing.T) {
	day := Day1{}

	var testDataPath, realDataPath string
	var err error

	if testDataPath, err = test.CreateTestDataPath("day1"); err != nil {
		t.Fatal(err)
	}
	if realDataPath, err = test.CreateRealDataPath("day1"); err != nil {
		t.Fatal(err)
	}

	t.Run("task 1", func(t *testing.T) {
		t.Run("test data", test.CreateTest(day.Task1, testDataPath, 24000))
		t.Run("real data", test.CreateTest(day.Task1, realDataPath, 67027))
	})

	t.Run("task 2", func(t *testing.T) {
		t.Run("test data", test.CreateTest(day.Task2, testDataPath, 45000))
		t.Run("real data", test.CreateTest(day.Task2, realDataPath, 197291))
	})
}
