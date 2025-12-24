package day9

import (
	tests "aoc/framework/tasks/tests"
	lo "github.com/samber/lo"
	"go/types"
	"testing"
)

func TestTask2(t *testing.T) {
	t.Run("mock(data1)", tests.RunTest(Task2, tests.MockData("data1"), lo.Empty[types.Nil](), 1, false))
	t.Run("mock(data2)", tests.RunTest(Task2, tests.MockData("data2"), lo.Empty[types.Nil](), 36, false))
	t.Run("real(day9)", tests.RunTest(Task2, tests.RealData("day9"), lo.Empty[types.Nil](), 2658, false))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("mock(data1)", tests.RunBenchmark(Task2, tests.MockData("data1"), lo.Empty[types.Nil](), false))
	b.Run("mock(data2)", tests.RunBenchmark(Task2, tests.MockData("data2"), lo.Empty[types.Nil](), false))
	b.Run("real(day9)", tests.RunBenchmark(Task2, tests.RealData("day9"), lo.Empty[types.Nil](), false))
}
