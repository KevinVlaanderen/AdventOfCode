package day9

import (
	"2023/src/test"
	"testing"
)

func TestTask1(t *testing.T) {
	var testDataPath, realDataPath string
	var err error

	if testDataPath, err = test.CreateTestDataPath("data"); err != nil {
		t.Fatal(err)
	}
	if realDataPath, err = test.CreateRealDataPath("day9"); err != nil {
		t.Fatal(err)
	}

	t.Run("test data", test.CreateTest(Task1, testDataPath, 114, false))
	t.Run("real data", test.CreateTest(Task1, realDataPath, 1993300041, false))
}

func TestTask2(t *testing.T) {
	var testDataPath, realDataPath string
	var err error

	if testDataPath, err = test.CreateTestDataPath("data"); err != nil {
		t.Fatal(err)
	}
	if realDataPath, err = test.CreateRealDataPath("day9"); err != nil {
		t.Fatal(err)
	}

	t.Run("test data", test.CreateTest(Task2, testDataPath, 2, false))
	t.Run("real data", test.CreateTest(Task2, realDataPath, 1038, false))
}
