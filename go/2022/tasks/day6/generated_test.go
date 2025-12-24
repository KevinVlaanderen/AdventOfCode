package day6

import (
	tests "aoc/framework/tasks/tests"
	lo "github.com/samber/lo"
	"go/types"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("mock(data1)", tests.RunTest(Task1, tests.MockData("data1"), lo.Empty[types.Nil](), 7, false))
	t.Run("mock(data2)", tests.RunTest(Task1, tests.MockData("data2"), lo.Empty[types.Nil](), 5, false))
	t.Run("mock(data3)", tests.RunTest(Task1, tests.MockData("data3"), lo.Empty[types.Nil](), 6, false))
	t.Run("mock(data4)", tests.RunTest(Task1, tests.MockData("data4"), lo.Empty[types.Nil](), 10, false))
	t.Run("mock(data5)", tests.RunTest(Task1, tests.MockData("data5"), lo.Empty[types.Nil](), 11, false))
	t.Run("real(day6)", tests.RunTest(Task1, tests.RealData("day6"), lo.Empty[types.Nil](), 1042, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("mock(data1)", tests.RunBenchmark(Task1, tests.MockData("data1"), lo.Empty[types.Nil](), false))
	b.Run("mock(data2)", tests.RunBenchmark(Task1, tests.MockData("data2"), lo.Empty[types.Nil](), false))
	b.Run("mock(data3)", tests.RunBenchmark(Task1, tests.MockData("data3"), lo.Empty[types.Nil](), false))
	b.Run("mock(data4)", tests.RunBenchmark(Task1, tests.MockData("data4"), lo.Empty[types.Nil](), false))
	b.Run("mock(data5)", tests.RunBenchmark(Task1, tests.MockData("data5"), lo.Empty[types.Nil](), false))
	b.Run("real(day6)", tests.RunBenchmark(Task1, tests.RealData("day6"), lo.Empty[types.Nil](), false))
}
func TestTask2(t *testing.T) {
	t.Run("mock(data1)", tests.RunTest(Task2, tests.MockData("data1"), lo.Empty[types.Nil](), 19, false))
	t.Run("mock(data2)", tests.RunTest(Task2, tests.MockData("data2"), lo.Empty[types.Nil](), 23, false))
	t.Run("mock(data3)", tests.RunTest(Task2, tests.MockData("data3"), lo.Empty[types.Nil](), 23, false))
	t.Run("mock(data4)", tests.RunTest(Task2, tests.MockData("data4"), lo.Empty[types.Nil](), 29, false))
	t.Run("mock(data5)", tests.RunTest(Task2, tests.MockData("data5"), lo.Empty[types.Nil](), 26, false))
	t.Run("real(day6)", tests.RunTest(Task2, tests.RealData("day6"), lo.Empty[types.Nil](), 2980, false))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("mock(data1)", tests.RunBenchmark(Task2, tests.MockData("data1"), lo.Empty[types.Nil](), false))
	b.Run("mock(data2)", tests.RunBenchmark(Task2, tests.MockData("data2"), lo.Empty[types.Nil](), false))
	b.Run("mock(data3)", tests.RunBenchmark(Task2, tests.MockData("data3"), lo.Empty[types.Nil](), false))
	b.Run("mock(data4)", tests.RunBenchmark(Task2, tests.MockData("data4"), lo.Empty[types.Nil](), false))
	b.Run("mock(data5)", tests.RunBenchmark(Task2, tests.MockData("data5"), lo.Empty[types.Nil](), false))
	b.Run("real(day6)", tests.RunBenchmark(Task2, tests.RealData("day6"), lo.Empty[types.Nil](), false))
}
