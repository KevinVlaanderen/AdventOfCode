package model

import (
	"log"
	"strconv"
	"strings"
)

type Instruction struct {
	Box         uint8
	Label       string
	Operation   Operation
	FocalLength uint8
}

func NewInstruction(data string, calculateHash func(word string) uint8) Instruction {
	signIndex := len(data) - 1
	operation := Remove
	var focalLength int

	if equalIndex := strings.IndexRune(data, '='); equalIndex >= 0 {
		signIndex = equalIndex
		operation = Set

		var err error
		if focalLength, err = strconv.Atoi(data[equalIndex+1:]); err != nil {
			log.Panicf("invalid number %v", data[signIndex+1:])
		}
	}

	label := data[:signIndex]
	return Instruction{
		Box:         calculateHash(label),
		Label:       label,
		Operation:   operation,
		FocalLength: uint8(focalLength),
	}
}

type Operation uint8

const (
	Remove Operation = iota
	Set
)
