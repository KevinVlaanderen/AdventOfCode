package day10

import (
	tests "aoc/framework/tasks/tests"
	lo "github.com/samber/lo"
	"go/types"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task1, tests.MockData("data"), lo.Empty[types.Nil](), 4, false))
	t.Run("Mock(data2)", tests.RunTest(Task1, tests.MockData("data2"), lo.Empty[types.Nil](), 8, false))
	t.Run("Real(day10)", tests.RunTest(Task1, tests.RealData("day10"), lo.Empty[types.Nil](), 6757, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("Mock(data2)", tests.RunBenchmark(Task1, tests.MockData("data2"), lo.Empty[types.Nil](), false))
	b.Run("Real(day10)", tests.RunBenchmark(Task1, tests.RealData("day10"), lo.Empty[types.Nil](), false))
}
func TestTask2(t *testing.T) {
	t.Run("Mock(data3)", tests.RunTest(Task2, tests.MockData("data3"), lo.Empty[types.Nil](), 4, false))
	t.Run("Mock(data4)", tests.RunTest(Task2, tests.MockData("data4"), lo.Empty[types.Nil](), 8, false))
	t.Run("Mock(data5)", tests.RunTest(Task2, tests.MockData("data5"), lo.Empty[types.Nil](), 10, false))
	t.Run("Real(day10)", tests.RunTest(Task2, tests.RealData("day10"), lo.Empty[types.Nil](), 523, false))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("Mock(data3)", tests.RunBenchmark(Task2, tests.MockData("data3"), lo.Empty[types.Nil](), false))
	b.Run("Mock(data4)", tests.RunBenchmark(Task2, tests.MockData("data4"), lo.Empty[types.Nil](), false))
	b.Run("Mock(data5)", tests.RunBenchmark(Task2, tests.MockData("data5"), lo.Empty[types.Nil](), false))
	b.Run("Real(day10)", tests.RunBenchmark(Task2, tests.RealData("day10"), lo.Empty[types.Nil](), false))
}
