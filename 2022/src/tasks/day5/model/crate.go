package model

import (
	"fmt"
	"regexp"
)

var cratePattern = regexp.MustCompile(`\[(\w+)]`)

type Crate struct {
	Label string
}

func NewCrate(data string) (crate Crate, err error) {
	matches := cratePattern.FindStringSubmatch(data)
	if len(matches) < 2 {
		return crate, fmt.Errorf("invalid crate data: %v", data)
	}
	return Crate{Label: matches[1]}, nil
}
