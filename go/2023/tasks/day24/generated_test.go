package day24

import (
	tests "aoc/framework/tasks/tests"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("mock(data)", tests.RunTest(Task1, tests.MockData("data"), "7,27", 2, false))
	t.Run("real(day24)", tests.RunTest(Task1, tests.RealData("day24"), "200000000000000,400000000000000", 21843, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), "7,27", false))
	b.Run("real(day24)", tests.RunBenchmark(Task1, tests.RealData("day24"), "200000000000000,400000000000000", false))
}
