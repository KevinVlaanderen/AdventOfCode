package day13

import (
	"2023/src/framework"
	"github.com/samber/lo"
)

func Task1(data string) (result framework.Result[int]) {
	blocks := framework.LineBlocks(data)

	horizontalIndices := make([]int, 0)
	verticalIndices := make([]int, 0)

Block:
	for _, block := range blocks {
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

		rowHashes := lo.Map(rows, func(row []bool, index int) framework.Hash {
			return framework.ComputeHash(row)
		})
		columnHashes := lo.Map(columns, func(column []bool, index int) framework.Hash {
			return framework.ComputeHash(column)
		})

	MidX:
		for mid := 0; mid < width-1; mid++ {
			checkN := min(mid+1, width-mid-1)
			for i := 0; i < checkN; i++ {
				if columnHashes[mid-i] != columnHashes[mid+i+1] {
					continue MidX
				}
			}
			verticalIndices = append(verticalIndices, mid+1)
			continue Block
		}

	MidY:
		for mid := 0; mid < height-1; mid++ {
			checkN := min(mid+1, height-mid-1)
			for i := 0; i < checkN; i++ {
				if rowHashes[mid-i] != rowHashes[mid+i+1] {
					continue MidY
				}
			}
			horizontalIndices = append(horizontalIndices, mid+1)
			continue Block
		}
	}

	for _, row := range horizontalIndices {
		result.Value += row * 100
	}
	for _, column := range verticalIndices {
		result.Value += column
	}

	return
}
