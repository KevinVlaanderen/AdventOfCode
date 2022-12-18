package day15

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
	if realDataPath, err = test.CreateRealDataPath("day15"); err != nil {
		t.Fatal(err)
	}

	t.Run("test data", test.CreateTest(CreateTask1(10), testDataPath, 26))
	t.Run("real data", test.CreateTest(CreateTask1(2000000), realDataPath, 4886370))
}
