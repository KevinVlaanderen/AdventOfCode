package day4

import (
	tests "aoc/framework/tasks/tests"
	lo "github.com/samber/lo"
	"go/types"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task1, tests.MockData("data"), lo.Empty[types.Nil](), 2, false))
	t.Run("Real(day4)", tests.RunTest(Task1, tests.RealData("day4"), lo.Empty[types.Nil](), 511, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("Real(day4)", tests.RunBenchmark(Task1, tests.RealData("day4"), lo.Empty[types.Nil](), false))
}
func TestTask2(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task2, tests.MockData("data"), lo.Empty[types.Nil](), 4, false))
	t.Run("Real(day4)", tests.RunTest(Task2, tests.RealData("day4"), lo.Empty[types.Nil](), 821, false))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("Real(day4)", tests.RunBenchmark(Task2, tests.RealData("day4"), lo.Empty[types.Nil](), false))
}
