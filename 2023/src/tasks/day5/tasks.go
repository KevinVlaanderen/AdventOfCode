package day5

import (
	"2023/src/framework/number"
	"2023/src/framework/parse"
	"2023/src/framework/task"
	"2023/src/tasks/day5/model"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"golang.org/x/exp/slices"
	"regexp"
)

var seedsPattern = regexp.MustCompile(`seeds: ([\d\s]+)`)

func Task1(filePath string) (result task.Result[int]) {
	blocks := task.Read(filePath, parse.Blocks())
	almanac := model.NewAlmanac(blocks[1:])

	seedsMatch := seedsPattern.FindStringSubmatch(blocks[0][0])
	seedNumbers := number.ExtractNumbers(seedsMatch[1])

	result.Value = lo.MinBy(lop.Map(seedNumbers, func(seed int, index int) int {
		return almanac.MapToValue(seed)
	}), func(a int, b int) bool {
		return a < b
	})

	return
}

func Task2(filePath string) (result task.Result[int]) {
	blocks := task.Read(filePath, parse.Blocks())
	almanac := model.NewAlmanac(blocks[1:])

	seedsMatch := seedsPattern.FindStringSubmatch(blocks[0][0])
	seedNumbers := number.ExtractNumbers(seedsMatch[1])
	seedRanges := lo.Map(lo.Chunk(seedNumbers, 2), func(item []int, index int) model.SeedRange {
		return model.NewSeedRange(item[0], item[1])
	})

	refinedSeedRanges := lo.FlatMap(seedRanges, func(seedRange model.SeedRange, index int) []model.SeedRange {
		return almanac.MapToRanges(seedRange)
	})

	slices.SortStableFunc(refinedSeedRanges, func(a, b model.SeedRange) int {
		return a.Start - b.Start
	})

	result.Value = refinedSeedRanges[0].Start
	return
}
