//go:build testing

package test

import (
	"testing"
)

func AssertEqual[Data comparable](t *testing.T, a Data, b Data) {
	if a == b {
		return
	}
	t.Errorf("Got %v, expected %v", a, b)
}
