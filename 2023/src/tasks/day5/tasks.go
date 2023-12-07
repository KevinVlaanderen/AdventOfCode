package day5

import (
	"2023/src/framework/generators"
	"2023/src/framework/number"
	"2023/src/framework/parse"
	"2023/src/framework/task"
	"2023/src/tasks/day5/model"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"regexp"
)

var seedsPattern = regexp.MustCompile(`seeds: ([\d\s]+)`)

func Task1(filePath string) (result task.Result[int]) {
	blocks := task.Read(filePath, parse.Blocks())
	almanac := model.NewAlmanac(blocks[1:])

	seedsMatch := seedsPattern.FindStringSubmatch(blocks[0][0])
	seedNumbers := number.ExtractNumbers(seedsMatch[1])
	seeds := seedNumbers

	location := -1
	for _, seed := range seeds {
		value := almanac.MapToNextStep("seed", seed)
		if location == -1 || value < location {
			location = value
		}
	}

	result.Value = location
	return
}

func Task2(filePath string) (result task.Result[int]) {
	blocks := task.Read(filePath, parse.Blocks())
	almanac := model.NewAlmanac(blocks[1:])

	seedsMatch := seedsPattern.FindStringSubmatch(blocks[0][0])
	seedNumbers := number.ExtractNumbers(seedsMatch[1])
	location := -1

	lop.ForEach(lo.Chunk(seedNumbers, 2), func(pair []int, pairIndex int) {
		for seed := range generators.RangeGen(pair[0], pair[1], 1) {
			value := almanac.MapToNextStep("seed", seed)
			if location == -1 || value < location {
				location = value
			}
		}
	})

	result.Value = location
	return
}
