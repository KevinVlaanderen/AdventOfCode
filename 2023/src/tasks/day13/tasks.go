package day13

import (
	"2023/src/framework"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

func Task1(data string) (result framework.Result[int]) {
	blocks := framework.LineBlocks(data)
	cache := framework.NewSafeCache[string, *HashGroup]()

	lop.ForEach(blocks, func(block []string, index int) {
		direction, index := findReflection(block, false, cache)
		switch direction {
		case Horizontal:
			result.Value += index * 100
		case Vertical:
			result.Value += index
		}
	})

	return
}

func Task2(data string) (result framework.Result[int]) {
	blocks := framework.LineBlocks(data)
	cache := framework.NewSafeCache[string, *HashGroup]()

	lop.ForEach(blocks, func(block []string, index int) {
		direction, index := findReflection(block, true, cache)
		switch direction {
		case Horizontal:
			result.Value += index * 100
		case Vertical:
			result.Value += index
		}
	})

	return
}

func findReflection(block []string, checkVariants bool, cache *framework.SafeCache[string, *HashGroup]) (Direction, int) {
	rows, columns := parseBlock(block)
	rowHashes := calculateHashes(rows, checkVariants, cache)
	columnHashes := calculateHashes(columns, checkVariants, cache)

	if found, index := reflectionFound(columnHashes, checkVariants); found {
		return Vertical, index
	}
	if found, index := reflectionFound(rowHashes, checkVariants); found {
		return Horizontal, index
	}
	panic("no reflection found")
}

func parseBlock(block []string) ([]string, []string) {
	width, height := len(block[0]), len(block)

	rows := make([]string, 0, height)
	columns := make([]string, width)

	for _, line := range block {
		rows = append(rows, line)
		for x, char := range line {
			columns[x] += string(char)
		}
	}
	return rows, columns
}

func calculateHashes(data []string, checkVariants bool, cache *framework.SafeCache[string, *HashGroup]) []*HashGroup {
	return lo.Map(data, func(item string, index int) *HashGroup {
		if group, found := cache.Get(item); found {
			return group
		} else {
			hashGroup := HashGroup{original: item}
			if checkVariants {
				hashGroup.variants = lo.Map([]rune(item), func(_ rune, index int) string {
					newRow := []rune(item)
					if item[index] == '#' {
						newRow[index] = '.'
					} else {
						newRow[index] = '#'
					}
					return string(newRow)
				})
			}
			cache.Set(item, &hashGroup)
			return &hashGroup
		}
	})
}

func reflectionFound(hashes []*HashGroup, checkVariants bool) (bool, int) {
	size := len(hashes)

	for mid := 0; mid < size-1; mid++ {
		checkN := min(mid+1, size-mid-1)

		matchFound := true
		variantUsed := false

		for offset := 0; offset < checkN; offset++ {
			if hashes[mid-offset].original == hashes[mid+offset+1].original {
				continue
			}
			if checkVariants && !variantUsed {
				if compareLeftRight(mid-offset, mid+offset+1, hashes) ||
					compareLeftRight(mid+offset+1, mid-offset, hashes) {
					variantUsed = true
					continue
				}
			}
			matchFound = false
		}
		if matchFound && (!checkVariants || variantUsed) {
			return true, mid + 1
		}
	}
	return false, 0
}

func compareLeftRight(originalIndex, variantIndex int, hashes []*HashGroup) bool {
	return lo.ContainsBy(hashes[variantIndex].variants, func(variant string) bool {
		return variant == hashes[originalIndex].original
	})
}

type HashGroup struct {
	original string
	variants []string
}

type Direction uint8

const (
	Horizontal Direction = iota
	Vertical
)
