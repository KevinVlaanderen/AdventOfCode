package day8

import (
	tests "aoc/framework/tasks/tests"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task1, tests.MockData("data"), 10, 40, false))
	t.Run("Real(day8)", tests.RunTest(Task1, tests.RealData("day8"), 1000, 122430, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), 10, false))
	b.Run("Real(day8)", tests.RunBenchmark(Task1, tests.RealData("day8"), 1000, false))
}
func TestTask2(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task2, tests.MockData("data"), 10, 25272, false))
	t.Run("Real(day8)", tests.RunTest(Task2, tests.RealData("day8"), 1000, 8135565324, false))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), 10, false))
	b.Run("Real(day8)", tests.RunBenchmark(Task2, tests.RealData("day8"), 1000, false))
}
