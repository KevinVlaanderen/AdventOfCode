package day6

import (
	"2022/src/test"
	"testing"
)

func TestTask1(t *testing.T) {
	var testDataPath1, testDataPath2, testDataPath3, testDataPath4, testDataPath5, realDataPath string
	var err error

	if testDataPath1, err = test.CreateTestDataPath("data1"); err != nil {
		t.Fatal(err)
	}
	if testDataPath2, err = test.CreateTestDataPath("data2"); err != nil {
		t.Fatal(err)
	}
	if testDataPath3, err = test.CreateTestDataPath("data3"); err != nil {
		t.Fatal(err)
	}
	if testDataPath4, err = test.CreateTestDataPath("data4"); err != nil {
		t.Fatal(err)
	}
	if testDataPath5, err = test.CreateTestDataPath("data5"); err != nil {
		t.Fatal(err)
	}

	if realDataPath, err = test.CreateRealDataPath("day6"); err != nil {
		t.Fatal(err)
	}

	t.Run("test data 1", test.CreateTest(Task1, testDataPath1, 7))
	t.Run("test data 2", test.CreateTest(Task1, testDataPath2, 5))
	t.Run("test data 3", test.CreateTest(Task1, testDataPath3, 6))
	t.Run("test data 4", test.CreateTest(Task1, testDataPath4, 10))
	t.Run("test data 5", test.CreateTest(Task1, testDataPath5, 11))
	t.Run("real data", test.CreateTest(Task1, realDataPath, 1042))
}

func TestTask2(t *testing.T) {
	var testDataPath1, testDataPath2, testDataPath3, testDataPath4, testDataPath5, realDataPath string
	var err error

	if testDataPath1, err = test.CreateTestDataPath("data1"); err != nil {
		t.Fatal(err)
	}
	if testDataPath2, err = test.CreateTestDataPath("data2"); err != nil {
		t.Fatal(err)
	}
	if testDataPath3, err = test.CreateTestDataPath("data3"); err != nil {
		t.Fatal(err)
	}
	if testDataPath4, err = test.CreateTestDataPath("data4"); err != nil {
		t.Fatal(err)
	}
	if testDataPath5, err = test.CreateTestDataPath("data5"); err != nil {
		t.Fatal(err)
	}

	if realDataPath, err = test.CreateRealDataPath("day6"); err != nil {
		t.Fatal(err)
	}

	t.Run("test data 1", test.CreateTest(Task2, testDataPath1, 19))
	t.Run("test data 2", test.CreateTest(Task2, testDataPath2, 23))
	t.Run("test data 3", test.CreateTest(Task2, testDataPath3, 23))
	t.Run("test data 4", test.CreateTest(Task2, testDataPath4, 29))
	t.Run("test data 5", test.CreateTest(Task2, testDataPath5, 26))
	t.Run("real data", test.CreateTest(Task2, realDataPath, 2980))
}
