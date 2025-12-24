package day11

import (
	tests "aoc/framework/tasks/tests"
	lo "github.com/samber/lo"
	"go/types"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("mock(data)", tests.RunTest(Task1, tests.MockData("data"), lo.Empty[types.Nil](), 10605, false))
	t.Run("real(day11)", tests.RunTest(Task1, tests.RealData("day11"), lo.Empty[types.Nil](), 76728, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("real(day11)", tests.RunBenchmark(Task1, tests.RealData("day11"), lo.Empty[types.Nil](), false))
}
