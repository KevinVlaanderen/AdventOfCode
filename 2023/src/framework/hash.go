package framework

import (
	"crypto/sha512"
	"fmt"
)

type Hash64 [64]byte

func ComputeHash64(data interface{}) Hash64 {
	return sha512.Sum512([]byte(fmt.Sprint(data)))
}
