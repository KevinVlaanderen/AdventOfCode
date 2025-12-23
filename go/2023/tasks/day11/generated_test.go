package day11

import (
	tests "aoc/framework/tasks/tests"
	lo "github.com/samber/lo"
	"go/types"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task1, tests.MockData("data"), lo.Empty[types.Nil](), 374, false))
	t.Run("Real(day11)", tests.RunTest(Task1, tests.RealData("day11"), lo.Empty[types.Nil](), 9418609, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("Real(day11)", tests.RunBenchmark(Task1, tests.RealData("day11"), lo.Empty[types.Nil](), false))
}
func TestTask2(t *testing.T) {
	t.Run("Real(day11)", tests.RunTest(Task2, tests.RealData("day11"), lo.Empty[types.Nil](), 593821230983, false))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("Real(day11)", tests.RunBenchmark(Task2, tests.RealData("day11"), lo.Empty[types.Nil](), false))
}
