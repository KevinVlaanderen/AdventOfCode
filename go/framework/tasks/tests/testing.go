package tests

import (
	"aoc/framework/tasks"
	"os"
	"strconv"
	"testing"
)

func RunTest[T comparable, P any](task tasks.Task[T, P], data DataDescriptor, param P, expected T, skipCI bool) func(*testing.T) {
	return func(t *testing.T) {
		if skipCI {
			if ci, ok := os.LookupEnv("CI"); ok {
				if ciBool, err := strconv.ParseBool(ci); err == nil && ciBool {
					t.SkipNow()
				}
			}
		}

		if d, err := ReadData(data); err != nil {
			t.Fatal(err)
			return
		} else {
			result := task(d, param)
			if result.Error != nil {
				t.Error(result.Error)
			} else {
				AssertEqual(t, result.Value, expected)
			}
		}
	}
}

func RunBenchmark[T comparable, P any](task tasks.Task[T, P], data DataDescriptor, param P, skipCI bool) func(*testing.B) {
	return func(b *testing.B) {
		if skipCI {
			if ci, ok := os.LookupEnv("CI"); ok {
				if ciBool, err := strconv.ParseBool(ci); err == nil && ciBool {
					b.SkipNow()
				}
			}
		}

		if d, err := ReadData(data); err != nil {
			b.Fatal(err)
			return
		} else {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				task(d, param)
			}
		}
	}
}
