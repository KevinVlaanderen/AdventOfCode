package day2

import (
	"2022/src/test"
	"testing"
)

func TestTask1(t *testing.T) {
	var testDataPath, realDataPath string
	var err error

	if testDataPath, err = test.CreateTestDataPath("data"); err != nil {
		t.Fatal(err)
	}
	if realDataPath, err = test.CreateRealDataPath("day2"); err != nil {
		t.Fatal(err)
	}

	t.Run("test data", test.CreateTest(Task1, testDataPath, 15))
	t.Run("real data", test.CreateTest(Task1, realDataPath, 11475))
}

func TestTask2(t *testing.T) {
	var testDataPath, realDataPath string
	var err error

	if testDataPath, err = test.CreateTestDataPath("data"); err != nil {
		t.Fatal(err)
	}
	if realDataPath, err = test.CreateRealDataPath("day2"); err != nil {
		t.Fatal(err)
	}

	t.Run("test data", test.CreateTest(Task2, testDataPath, 12))
	t.Run("real data", test.CreateTest(Task2, realDataPath, 16862))
}
