package day19

import (
	"2023/src/framework"
	"2023/src/tasks/day19/model"
	"github.com/samber/lo"
	"go/types"
)

func Task1(data string, _ types.Nil) (result framework.Result[int]) {
	blocks := framework.Blocks(data)
	workflows, parts := model.ParseWorkflows(blocks[0]), model.ParseParts(blocks[1])
	system := model.NewSystem(workflows)

	validBoundaries := system.ValidBoundaries("in", "A")

	result.Value = lo.SumBy(parts, func(part model.Part) int {
		if Check(validBoundaries, part) {
			return part.Rating()
		}
		return 0
	})

	//system.Draw("system.gv")

	return
}

func Task2(data string, _ types.Nil) (result framework.Result[int]) {
	blocks := framework.Blocks(data)
	workflows := model.ParseWorkflows(blocks[0])
	system := model.NewSystem(workflows)

	validBoundaries := system.ValidBoundaries("in", "A")

	total, overlap := 0, 0
	keys := lo.Keys(validBoundaries[0])

	for _, boundary := range validBoundaries {
		total += lo.Reduce(lo.Values(boundary), func(agg int, value lo.Tuple2[int, int], index int) int {
			return agg * (value.B - value.A + 1)
		}, 1)
	}

	overlaps := make([]int, len(keys))
	for i := 0; i < len(validBoundaries)-1; i++ {
		for j := i + 1; j < len(validBoundaries); j++ {
			for index, key := range keys {
				overlaps[index] = 0
				first, second := validBoundaries[i][key], validBoundaries[j][key]
				if first.B >= second.A && first.A <= second.B {
					overlaps[index] = first.B - second.A + 1
				}
			}

			overlap += lo.Reduce(overlaps, func(agg int, overlap int, index int) int {
				return agg * overlap
			}, 1)
		}
	}

	result.Value = total - overlap

	return
}

func Check(validBoundaries []map[model.Category]lo.Tuple2[int, int], part model.Part) bool {
Boundary:
	for _, validBoundary := range validBoundaries {
		for category, value := range part.Ratings {
			if value < validBoundary[category].A || value > validBoundary[category].B {
				continue Boundary
			}
		}
		return true
	}

	return false
}
