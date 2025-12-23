package day18

import (
	tests "aoc/framework/tasks/tests"
	lo "github.com/samber/lo"
	"go/types"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task1, tests.MockData("data"), lo.Empty[types.Nil](), 62, false))
	t.Run("Real(day18)", tests.RunTest(Task1, tests.RealData("day18"), lo.Empty[types.Nil](), 70026, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("Real(day18)", tests.RunBenchmark(Task1, tests.RealData("day18"), lo.Empty[types.Nil](), false))
}
func TestTask2(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task2, tests.MockData("data"), lo.Empty[types.Nil](), 952408144115, false))
	t.Run("Real(day18)", tests.RunTest(Task2, tests.RealData("day18"), lo.Empty[types.Nil](), 68548301037382, false))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("Real(day18)", tests.RunBenchmark(Task2, tests.RealData("day18"), lo.Empty[types.Nil](), false))
}
