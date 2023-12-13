package framework

import (
	"crypto/sha512"
	"fmt"
)

type Hash [64]byte

func ComputeHash(data interface{}) Hash {
	return sha512.Sum512([]byte(fmt.Sprint(data)))
}
