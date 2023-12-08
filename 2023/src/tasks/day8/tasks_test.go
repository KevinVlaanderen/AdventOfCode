package day8

import (
	"2023/src/test"
	"testing"
)

func TestTask1(t *testing.T) {
	var testDataPath1, testDataPath2, realDataPath string
	var err error

	if testDataPath1, err = test.CreateTestDataPath("data"); err != nil {
		t.Fatal(err)
	}
	if testDataPath2, err = test.CreateTestDataPath("data2"); err != nil {
		t.Fatal(err)
	}
	if realDataPath, err = test.CreateRealDataPath("day8"); err != nil {
		t.Fatal(err)
	}

	t.Run("test data", test.CreateTest(Task1, testDataPath1, 2, false))
	t.Run("test data", test.CreateTest(Task1, testDataPath2, 6, false))
	t.Run("real data", test.CreateTest(Task1, realDataPath, 14893, false))
}
