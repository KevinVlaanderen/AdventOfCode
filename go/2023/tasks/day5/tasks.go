//go:generate go run aoc/cmd/generate-tests

package day5

import (
	"aoc/2023/tasks/day5/model"
	"aoc/framework/math"
	"aoc/framework/tasks"
	"go/types"
	"regexp"

	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

var seedsPattern = regexp.MustCompile(`seeds: ([\d\s]+)`)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	blocks := tasks.LineBlocks(data)
	almanac := model.NewAlmanac(blocks[1:])

	seedsMatch := seedsPattern.FindStringSubmatch(blocks[0][0])
	seedNumbers := math.ExtractNumbers(seedsMatch[1])

	result.Value = lo.MinBy(lo.Map(seedNumbers, func(seed int, index int) int {
		return almanac.MapToValue(seed)
	}), func(a int, b int) bool {
		return a < b
	})

	return
}

func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	blocks := tasks.LineBlocks(data)
	almanac := model.NewAlmanac(blocks[1:])

	seedsMatch := seedsPattern.FindStringSubmatch(blocks[0][0])
	seedNumbers := math.ExtractNumbers(seedsMatch[1])
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
