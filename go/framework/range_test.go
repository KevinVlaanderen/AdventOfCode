package framework

import (
	"runtime"
	"strconv"
	"testing"
)

var sizes = []int{10, 10000, 10000000}

func BenchmarkRangeSlice(b *testing.B) {
	for _, size := range sizes {
		b.Run(strconv.Itoa(size), func(b *testing.B) {
			b.ReportAllocs()

			var result int

			for b.Loop() {
				for v := range RangeSlice(0, size, 1) {
					result += v
				}
			}

			runtime.KeepAlive(result)
		})
	}
}

func BenchmarkRangeChan(b *testing.B) {
	for _, size := range sizes {
		b.Run(strconv.Itoa(size), func(b *testing.B) {
			b.ReportAllocs()

			var result int

			for b.Loop() {
				for v := range RangeChan(0, size, 1) {
					result += v
				}
			}

			runtime.KeepAlive(result)
		})
	}
}

func BenchmarkRangeYield(b *testing.B) {
	for _, size := range sizes {
		b.Run(strconv.Itoa(size), func(b *testing.B) {
			b.ReportAllocs()

			var result int

			for b.Loop() {
				for v := range RangeYield(0, size, 1) {
					result += v
				}
			}

			runtime.KeepAlive(result)
		})
	}
}
