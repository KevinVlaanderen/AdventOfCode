package day12

import (
	tests "aoc/framework/tasks/tests"
	lo "github.com/samber/lo"
	"go/types"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("mock(data)", tests.RunTest(Task1, tests.MockData("data"), lo.Empty[types.Nil](), 31, false))
	t.Run("real(day12)", tests.RunTest(Task1, tests.RealData("day12"), lo.Empty[types.Nil](), 330, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("real(day12)", tests.RunBenchmark(Task1, tests.RealData("day12"), lo.Empty[types.Nil](), false))
}
func TestTask2(t *testing.T) {
	t.Run("mock(data)", tests.RunTest(Task2, tests.MockData("data"), lo.Empty[types.Nil](), 29, false))
	t.Run("real(day12)", tests.RunTest(Task2, tests.RealData("day12"), lo.Empty[types.Nil](), 321, false))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("real(day12)", tests.RunBenchmark(Task2, tests.RealData("day12"), lo.Empty[types.Nil](), false))
}
