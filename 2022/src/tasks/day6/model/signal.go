package model

import "math/bits"

type Signal string

func (s Signal) FindMarker(length int) int {
	for i := 0; i < len(s)-length; i++ {
		slice := s[i : i+length]

		var unique uint32
		for _, char := range slice {
			unique |= 1 << (int(char) - 97)
		}
		if bits.OnesCount32(unique) == length {
			return i + length
		}
	}
	return -1
}
