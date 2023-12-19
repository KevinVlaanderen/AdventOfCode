package day19

import (
	"2023/src/framework"
	"2023/src/tasks/day19/model"
	"github.com/samber/lo"
)

func Task1(data string) (result framework.Result[int]) {
	system, parts := parse(data)

	result.Value = lo.SumBy(parts, func(part model.Part) int {
		if system.Check(part) {
			return part.Rating()
		}
		return 0
	})

	return
}

func parse(data string) (model.System, []model.Part) {
	blocks := framework.Blocks(data)
	return model.ParseSystem(blocks[0]), model.ParseParts(blocks[1])
}
