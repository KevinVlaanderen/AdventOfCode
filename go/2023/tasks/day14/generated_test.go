package day14

import (
	tests "aoc/framework/tasks/tests"
	lo "github.com/samber/lo"
	"go/types"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task1, tests.MockData("data"), lo.Empty[types.Nil](), 136, false))
	t.Run("Real(day14)", tests.RunTest(Task1, tests.RealData("day14"), lo.Empty[types.Nil](), 113525, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("Real(day14)", tests.RunBenchmark(Task1, tests.RealData("day14"), lo.Empty[types.Nil](), false))
}
func TestTask2(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task2, tests.MockData("data"), lo.Empty[types.Nil](), 64, false))
	t.Run("Real(day14)", tests.RunTest(Task2, tests.RealData("day14"), lo.Empty[types.Nil](), 101292, false))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("Real(day14)", tests.RunBenchmark(Task2, tests.RealData("day14"), lo.Empty[types.Nil](), false))
}
