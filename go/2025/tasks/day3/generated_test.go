package day3

import (
	tests "aoc/framework/tasks/tests"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task1, tests.MockData("data"), 2, 357, false))
	t.Run("Real(day3)", tests.RunTest(Task1, tests.RealData("day3"), 2, 17109, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), 2, false))
	b.Run("Real(day3)", tests.RunBenchmark(Task1, tests.RealData("day3"), 2, false))
}
func TestTask2(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task2, tests.MockData("data"), 12, 3121910778619, false))
	t.Run("Real(day3)", tests.RunTest(Task2, tests.RealData("day3"), 12, 169347417057382, false))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), 12, false))
	b.Run("Real(day3)", tests.RunBenchmark(Task2, tests.RealData("day3"), 12, false))
}
