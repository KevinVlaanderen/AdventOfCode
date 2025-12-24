package day9

import (
	tests "aoc/framework/tasks/tests"
	lo "github.com/samber/lo"
	"go/types"
	"testing"
)

func TestTask2(t *testing.T) {
	t.Run("mock(data)", tests.RunTest(Task2, tests.MockData("data"), lo.Empty[types.Nil](), 24, false))
	t.Run("real(day9)", tests.RunTest(Task2, tests.RealData("day9"), lo.Empty[types.Nil](), 1473551379, false))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("real(day9)", tests.RunBenchmark(Task2, tests.RealData("day9"), lo.Empty[types.Nil](), false))
}
func TestTask1(t *testing.T) {
	t.Run("mock(data)", tests.RunTest(Task1, tests.MockData("data"), lo.Empty[types.Nil](), 50, false))
	t.Run("real(day9)", tests.RunTest(Task1, tests.RealData("day9"), lo.Empty[types.Nil](), 4777409595, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("real(day9)", tests.RunBenchmark(Task1, tests.RealData("day9"), lo.Empty[types.Nil](), false))
}
