package day3

import (
	"2022/src/framework"
	"fmt"
	"math/bits"
	"unicode"
)

func Task1(filePath string) (result framework.TaskResult) {
	data, err := framework.ReadLines(filePath)
	if err != nil {
		return framework.TaskResult{Error: err}
	}

	for _, line := range data {
		rucksack := NewRucksack(line)
		if difference, err := Intersection([][]rune{rucksack.compartment1, rucksack.compartment2}); err != nil {
			return framework.TaskResult{Error: err}
		} else {
			if len(difference) != 1 {
				return framework.TaskResult{Error: fmt.Errorf("invalid number of differences in %s", line)}
			}

			result.Value += GetPriority(difference[0])
		}
	}

	return
}

func Task2(filePath string) (result framework.TaskResult) {
	_, err := framework.ReadLines(filePath)
	if err != nil {
		return framework.TaskResult{Error: err}
	}

	return
}

type Rucksack struct {
	compartment1 []rune
	compartment2 []rune
}

func NewRucksack(content string) Rucksack {
	runes := []rune(content)
	return Rucksack{
		compartment1: runes[:len(runes)/2],
		compartment2: runes[len(runes)/2:],
	}
}

func Intersection[T comparable](slices [][]T) (intersection []T, err error) {
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

func GetPriority(item rune) int {
	if unicode.IsLower(item) {
		return int(item) - 96
	} else {
		return int(item) - 38
	}
}
