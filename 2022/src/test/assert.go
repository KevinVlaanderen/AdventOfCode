//go:build testing

package test

import (
	"testing"
)

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Errorf("Got %v, expected %v", a, b)
}
