package day15

import (
	tests "aoc/framework/tasks/tests"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("mock(data)", tests.RunTest(Task1, tests.MockData("data"), 10, 26, false))
	t.Run("real(day15)", tests.RunTest(Task1, tests.RealData("day15"), 2000000, 4886370, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), 10, false))
	b.Run("real(day15)", tests.RunBenchmark(Task1, tests.RealData("day15"), 2000000, false))
}
