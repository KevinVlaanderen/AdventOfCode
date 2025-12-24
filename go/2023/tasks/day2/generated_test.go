package day2

import (
	tests "aoc/framework/tasks/tests"
	lo "github.com/samber/lo"
	"go/types"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("mock(data)", tests.RunTest(Task1, tests.MockData("data"), lo.Empty[types.Nil](), 8, false))
	t.Run("real(day2)", tests.RunTest(Task1, tests.RealData("day2"), lo.Empty[types.Nil](), 2278, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("real(day2)", tests.RunBenchmark(Task1, tests.RealData("day2"), lo.Empty[types.Nil](), false))
}
func TestTask2(t *testing.T) {
	t.Run("mock(data)", tests.RunTest(Task2, tests.MockData("data"), lo.Empty[types.Nil](), 2286, false))
	t.Run("real(day2)", tests.RunTest(Task2, tests.RealData("day2"), lo.Empty[types.Nil](), 67953, false))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("real(day2)", tests.RunBenchmark(Task2, tests.RealData("day2"), lo.Empty[types.Nil](), false))
}
