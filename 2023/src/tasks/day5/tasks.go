package day5

import (
	"2023/src/framework/parse"
	"2023/src/framework/task"
	"2023/src/tasks/day5/model"
	"golang.org/x/exp/slices"
)

func Task1(filePath string) (result task.Result[int]) {
	blocks := task.Read(filePath, parse.Blocks())
	almanac := model.NewAlmanac(blocks)

	locations := make([]int, len(almanac.Seeds))
	for index, seed := range almanac.Seeds {
		locations[index] = almanac.MapToNext("seed", seed)
	}

	result.Value = slices.Min(locations)
	return
}
