package model

import "unicode"

type Rucksack []Item

func (r Rucksack) Compartment1() []Item {
	length := len(r)
	return r[:length/2]
}

func (r Rucksack) Compartment2() []Item {
	length := len(r)
	return r[length/2:]
}

type Item rune

func (i Item) GetPriority() int {
	if unicode.IsLower(rune(i)) {
		return int(i) - 96
	} else {
		return int(i) - 38
	}
}
