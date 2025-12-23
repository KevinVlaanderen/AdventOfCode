package day1

import (
	tests "aoc/framework/tasks/tests"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task1, tests.MockData("data"), false, 3, false))
	t.Run("Real(day1)", tests.RunTest(Task1, tests.RealData("day1"), false, 1023, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), false, false))
	b.Run("Real(day1)", tests.RunBenchmark(Task1, tests.RealData("day1"), false, false))
}
func TestTask2(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task2, tests.MockData("data"), true, 6, false))
	t.Run("Real(day1)", tests.RunTest(Task2, tests.RealData("day1"), true, 5899, false))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), true, false))
	b.Run("Real(day1)", tests.RunBenchmark(Task2, tests.RealData("day1"), true, false))
}
