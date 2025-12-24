package day21

import (
	tests "aoc/framework/tasks/tests"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("mock(data)", tests.RunTest(Task1, tests.MockData("data"), 6, 16, false))
	t.Run("real(day21)", tests.RunTest(Task1, tests.RealData("day21"), 64, 3646, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), 6, false))
	b.Run("real(day21)", tests.RunBenchmark(Task1, tests.RealData("day21"), 64, false))
}
func TestTask2(t *testing.T) {
	t.Run("mock(data)", tests.RunTest(Task2, tests.MockData("data"), 6, 16, true))
	t.Run("mock(data)", tests.RunTest(Task2, tests.MockData("data"), 10, 50, true))
	t.Run("mock(data)", tests.RunTest(Task2, tests.MockData("data"), 50, 1594, true))
	t.Run("mock(data)", tests.RunTest(Task2, tests.MockData("data"), 100, 6536, true))
	t.Run("mock(data)", tests.RunTest(Task2, tests.MockData("data"), 500, 167004, true))
	t.Run("mock(data)", tests.RunTest(Task2, tests.MockData("data"), 1000, 668697, true))
	t.Run("mock(data)", tests.RunTest(Task2, tests.MockData("data"), 5000, 16733044, true))
	t.Run("real(day21)", tests.RunTest(Task2, tests.RealData("day21"), 26501365, -1, true))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), 6, true))
	b.Run("mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), 10, true))
	b.Run("mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), 50, true))
	b.Run("mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), 100, true))
	b.Run("mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), 500, true))
	b.Run("mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), 1000, true))
	b.Run("mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), 5000, true))
	b.Run("real(day21)", tests.RunBenchmark(Task2, tests.RealData("day21"), 26501365, true))
}
