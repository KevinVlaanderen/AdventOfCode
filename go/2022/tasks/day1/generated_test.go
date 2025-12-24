package day1

import (
	tests "aoc/framework/tasks/tests"
	lo "github.com/samber/lo"
	"go/types"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("mock(data)", tests.RunTest(Task1, tests.MockData("data"), lo.Empty[types.Nil](), 24000, false))
	t.Run("real(day1)", tests.RunTest(Task1, tests.RealData("day1"), lo.Empty[types.Nil](), 67027, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("real(day1)", tests.RunBenchmark(Task1, tests.RealData("day1"), lo.Empty[types.Nil](), false))
}
func TestTask2(t *testing.T) {
	t.Run("mock(data)", tests.RunTest(Task2, tests.MockData("data"), lo.Empty[types.Nil](), 45000, false))
	t.Run("real(day1)", tests.RunTest(Task2, tests.RealData("day1"), lo.Empty[types.Nil](), 197291, false))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("real(day1)", tests.RunBenchmark(Task2, tests.RealData("day1"), lo.Empty[types.Nil](), false))
}
