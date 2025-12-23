package day16

import (
	tests "aoc/framework/tasks/tests"
	lo "github.com/samber/lo"
	"go/types"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task1, tests.MockData("data"), lo.Empty[types.Nil](), 46, false))
	t.Run("Real(day16)", tests.RunTest(Task1, tests.RealData("day16"), lo.Empty[types.Nil](), 7210, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("Real(day16)", tests.RunBenchmark(Task1, tests.RealData("day16"), lo.Empty[types.Nil](), false))
}
func TestTask2(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task2, tests.MockData("data"), lo.Empty[types.Nil](), 51, false))
	t.Run("Real(day16)", tests.RunTest(Task2, tests.RealData("day16"), lo.Empty[types.Nil](), 7673, false))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("Real(day16)", tests.RunBenchmark(Task2, tests.RealData("day16"), lo.Empty[types.Nil](), false))
}
