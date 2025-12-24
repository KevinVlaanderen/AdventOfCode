package day20

import (
	tests "aoc/framework/tasks/tests"
	lo "github.com/samber/lo"
	"go/types"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("mock(data)", tests.RunTest(Task1, tests.MockData("data"), lo.Empty[types.Nil](), 32000000, false))
	t.Run("mock(data2)", tests.RunTest(Task1, tests.MockData("data2"), lo.Empty[types.Nil](), 11687500, false))
	t.Run("real(day20)", tests.RunTest(Task1, tests.RealData("day20"), lo.Empty[types.Nil](), 821985143, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("mock(data2)", tests.RunBenchmark(Task1, tests.MockData("data2"), lo.Empty[types.Nil](), false))
	b.Run("real(day20)", tests.RunBenchmark(Task1, tests.RealData("day20"), lo.Empty[types.Nil](), false))
}
func TestTask2(t *testing.T) {
	t.Run("real(day20)", tests.RunTest(Task2, tests.RealData("day20"), lo.Empty[types.Nil](), -1, true))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("real(day20)", tests.RunBenchmark(Task2, tests.RealData("day20"), lo.Empty[types.Nil](), true))
}
