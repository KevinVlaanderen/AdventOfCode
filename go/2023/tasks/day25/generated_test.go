package day25

import (
	tests "aoc/framework/tasks/tests"
	lo "github.com/samber/lo"
	"go/types"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("mock(data)", tests.RunTest(Task1, tests.MockData("data"), lo.Empty[types.Nil](), 54, false))
	t.Run("real(day25)", tests.RunTest(Task1, tests.RealData("day25"), lo.Empty[types.Nil](), -1, true))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("real(day25)", tests.RunBenchmark(Task1, tests.RealData("day25"), lo.Empty[types.Nil](), true))
}
