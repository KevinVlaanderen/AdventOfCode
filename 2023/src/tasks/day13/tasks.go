package day13

import (
	"2023/src/framework"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

func Task1(data string) (result framework.Result[int]) {
	blocks := framework.LineBlocks(data)

	lop.ForEach(blocks, func(block []string, index int) {
		direction, index := findReflection(block, false)
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

	lop.ForEach(blocks, func(block []string, index int) {
		direction, index := findReflection(block, true)
		switch direction {
		case Horizontal:
			result.Value += index * 100
		case Vertical:
			result.Value += index
		}
	})

	return
}

func findReflection(block []string, checkVariants bool) (Direction, int) {
	rows, columns := parseBlock(block)

	rowHashes := calculateHashes(rows, checkVariants)
	columnHashes := calculateHashes(columns, checkVariants)

	if found, index := reflectionFound(columnHashes, checkVariants); found {
		return Vertical, index
	}
	if found, index := reflectionFound(rowHashes, checkVariants); found {
		return Horizontal, index
	}
	panic("no reflection found")
}

func parseBlock(block []string) ([][]bool, [][]bool) {
	width, height := len(block[0]), len(block)

	rows := make([][]bool, height)
	columns := make([][]bool, width)

	for y, line := range block {
		if rows[y] == nil {
			rows[y] = make([]bool, width)
		}
		for x, char := range line {
			if columns[x] == nil {
				columns[x] = make([]bool, height)
			}
			rows[y][x] = char == '#'
			columns[x][y] = char == '#'
		}
	}
	return rows, columns
}

func calculateHashes(data [][]bool, checkVariants bool) []HashGroup {
	return lo.Map(data, func(item []bool, index int) (hashGroup HashGroup) {
		hashGroup.original = framework.ComputeHash(item)
		if checkVariants {
			hashGroup.variants = lo.Map(item, func(_ bool, index int) framework.Hash {
				newRow := make([]bool, len(item))
				copy(newRow, item)
				newRow[index] = !item[index]
				return framework.ComputeHash(newRow)
			})
		}
		return
	})
}

func reflectionFound(hashes []HashGroup, checkVariants bool) (bool, int) {
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

func compareLeftRight(originalIndex, variantIndex int, hashes []HashGroup) bool {
	return lo.ContainsBy(hashes[variantIndex].variants, func(variant framework.Hash) bool {
		return variant == hashes[originalIndex].original
	})
}

type HashGroup struct {
	original framework.Hash
	variants []framework.Hash
}

type Direction uint8

const (
	Horizontal Direction = iota
	Vertical
)
