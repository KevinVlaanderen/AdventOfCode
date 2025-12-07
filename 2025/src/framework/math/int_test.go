package math

import "testing"

func BenchmarkLength(b *testing.B) {
	for b.Loop() {
		_ = Length(123456789)
	}
}

func BenchmarkLength2(b *testing.B) {
	for b.Loop() {
		_, _ = Length2(123456789)
	}
}
