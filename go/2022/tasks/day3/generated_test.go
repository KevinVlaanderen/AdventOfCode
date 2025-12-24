package day3

import (
	tests "aoc/framework/tasks/tests"
	lo "github.com/samber/lo"
	"go/types"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("mock(data)", tests.RunTest(Task1, tests.MockData("data"), lo.Empty[types.Nil](), 157, false))
	t.Run("real(day3)", tests.RunTest(Task1, tests.RealData("day3"), lo.Empty[types.Nil](), 8252, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("real(day3)", tests.RunBenchmark(Task1, tests.RealData("day3"), lo.Empty[types.Nil](), false))
}
func TestTask2(t *testing.T) {
	t.Run("mock(data)", tests.RunTest(Task2, tests.MockData("data"), lo.Empty[types.Nil](), 70, false))
	t.Run("real(day3)", tests.RunTest(Task2, tests.RealData("day3"), lo.Empty[types.Nil](), 2828, false))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("real(day3)", tests.RunBenchmark(Task2, tests.RealData("day3"), lo.Empty[types.Nil](), false))
}
