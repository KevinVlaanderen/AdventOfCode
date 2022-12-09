package day9

import (
	"2022/src/test"
	"testing"
)

func TestTask1(t *testing.T) {
	var testDataPath1, realDataPath string
	var err error

	if testDataPath1, err = test.CreateTestDataPath("data1"); err != nil {
		t.Fatal(err)
	}
	if realDataPath, err = test.CreateRealDataPath("day9"); err != nil {
		t.Fatal(err)
	}

	t.Run("test data 1", test.CreateTest(Task1, testDataPath1, 13))
	t.Run("real data", test.CreateTest(Task1, realDataPath, 6470))
}

func TestTask2(t *testing.T) {
	var testDataPath1, testDataPath2, realDataPath string
	var err error

	if testDataPath1, err = test.CreateTestDataPath("data1"); err != nil {
		t.Fatal(err)
	}
	if testDataPath2, err = test.CreateTestDataPath("data2"); err != nil {
		t.Fatal(err)
	}
	if realDataPath, err = test.CreateRealDataPath("day9"); err != nil {
		t.Fatal(err)
	}

	t.Run("test data 1", test.CreateTest(Task2, testDataPath1, 1))
	t.Run("test data 2", test.CreateTest(Task2, testDataPath2, 36))
	t.Run("real data", test.CreateTest(Task2, realDataPath, 2658))
}
