package day15

import (
	tests "aoc/framework/tasks/tests"
	lo "github.com/samber/lo"
	"go/types"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task1, tests.MockData("data"), lo.Empty[types.Nil](), 52, false))
	t.Run("Mock(data2)", tests.RunTest(Task1, tests.MockData("data2"), lo.Empty[types.Nil](), 1320, false))
	t.Run("Real(day15)", tests.RunTest(Task1, tests.RealData("day15"), lo.Empty[types.Nil](), 505427, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("Mock(data2)", tests.RunBenchmark(Task1, tests.MockData("data2"), lo.Empty[types.Nil](), false))
	b.Run("Real(day15)", tests.RunBenchmark(Task1, tests.RealData("day15"), lo.Empty[types.Nil](), false))
}
func TestTask2(t *testing.T) {
	t.Run("Mock(data2)", tests.RunTest(Task2, tests.MockData("data2"), lo.Empty[types.Nil](), 145, false))
	t.Run("Real(day15)", tests.RunTest(Task2, tests.RealData("day15"), lo.Empty[types.Nil](), 243747, false))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("Mock(data2)", tests.RunBenchmark(Task2, tests.MockData("data2"), lo.Empty[types.Nil](), false))
	b.Run("Real(day15)", tests.RunBenchmark(Task2, tests.RealData("day15"), lo.Empty[types.Nil](), false))
}
