package framework

import (
	"fmt"
	"math/bits"
)

func Intersection[T comparable](slices ...[]T) (intersection []T, err error) {
	num := len(slices)
	if num > 64 {
		return []T{}, fmt.Errorf("too many slices provides: %d", num)
	}

	m := make(map[T]uint64)

	for index, slice := range slices {
		for _, k := range slice {
			m[k] |= 1 << index
		}
	}

	for k, v := range m {
		if bits.OnesCount64(v) == num {
			intersection = append(intersection, k)
		}
	}

	return
}
