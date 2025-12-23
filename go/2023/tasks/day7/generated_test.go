package day7

import (
	tests "aoc/framework/tasks/tests"
	lo "github.com/samber/lo"
	"go/types"
	"testing"
)

func TestTask1(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task1, tests.MockData("data"), lo.Empty[types.Nil](), 6440, false))
	t.Run("Real(day7)", tests.RunTest(Task1, tests.RealData("day7"), lo.Empty[types.Nil](), 251216224, false))
}
func BenchmarkTask1(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task1, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("Real(day7)", tests.RunBenchmark(Task1, tests.RealData("day7"), lo.Empty[types.Nil](), false))
}
func TestTask2(t *testing.T) {
	t.Run("Mock(data)", tests.RunTest(Task2, tests.MockData("data"), lo.Empty[types.Nil](), 5905, false))
	t.Run("Real(day7)", tests.RunTest(Task2, tests.RealData("day7"), lo.Empty[types.Nil](), 250825971, false))
}
func BenchmarkTask2(b *testing.B) {
	b.Run("Mock(data)", tests.RunBenchmark(Task2, tests.MockData("data"), lo.Empty[types.Nil](), false))
	b.Run("Real(day7)", tests.RunBenchmark(Task2, tests.RealData("day7"), lo.Empty[types.Nil](), false))
}
