package framework

import (
	"bytes"
	"crypto/sha512"
	"fmt"
	"strings"
)

type Hash64 [64]byte

func ComputeHash64(data interface{}) Hash64 {
	return sha512.Sum512([]byte(fmt.Sprint(data)))
}

func ComputeSliceHash64[T any](list []T) Hash64 {
	var buffer bytes.Buffer
	for i := range list {
		buffer.WriteString(fmt.Sprintf("%v,", list[i]))
	}
	return sha512.Sum512([]byte(buffer.String()))
}

func ComputeStringSliceHash64(list []string) Hash64 {
	return sha512.Sum512([]byte(strings.Join(list, "")))
}
