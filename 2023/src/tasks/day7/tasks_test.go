package day7

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
	if realDataPath, err = test.CreateRealDataPath("day7"); err != nil {
		t.Fatal(err)
	}

	t.Run("test data", test.CreateTest(Task1, testDataPath, 6440))
	t.Run("real data", test.CreateTest(Task1, realDataPath, 251216224))
}

func TestTask2(t *testing.T) {
	var testDataPath, realDataPath string
	var err error

	if testDataPath, err = test.CreateTestDataPath("data"); err != nil {
		t.Fatal(err)
	}
	if realDataPath, err = test.CreateRealDataPath("day7"); err != nil {
		t.Fatal(err)
	}

	t.Run("test data", test.CreateTest(Task2, testDataPath, 5905))
	t.Run("real data", test.CreateTest(Task2, realDataPath, 250825971))
}
