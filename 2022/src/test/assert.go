//go:build testing

package test

import (
	"fmt"
	"testing"
)

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Fatal(fmt.Sprintf("Got %v, expected %v", a, b))
}
