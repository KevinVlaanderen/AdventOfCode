package day5

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
	if realDataPath, err = test.CreateRealDataPath("day5"); err != nil {
		t.Fatal(err)
	}

	t.Run("test data", test.CreateTest(Task1, testDataPath, "CMZ"))
	t.Run("real data", test.CreateTest(Task1, realDataPath, "SHQWSRBDL"))
}

func TestTask2(t *testing.T) {
	var testDataPath, realDataPath string
	var err error

	if testDataPath, err = test.CreateTestDataPath("data"); err != nil {
		t.Fatal(err)
	}
	if realDataPath, err = test.CreateRealDataPath("day5"); err != nil {
		t.Fatal(err)
	}

	t.Run("test data", test.CreateTest(Task2, testDataPath, "MCD"))
	t.Run("real data", test.CreateTest(Task2, realDataPath, "CDTQZHBRS"))
}
