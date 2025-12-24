package day5

import (
	"aoc/framework/tasks"
	"go/types"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Ingredients struct {
	FreshRanges []*Range
	Ingredients []int
}

type Range struct {
	From, To int
}

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	ingredients := parse(data)

	for _, ingredient := range ingredients.Ingredients {
		for _, freshRange := range ingredients.FreshRanges {
			if ingredient >= freshRange.From && ingredient <= freshRange.To {
				result.Value++
				break
			}
		}
	}

	return
}

func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	ingredients := parse(data)

	mergedRanges := mergeRanges(ingredients.FreshRanges)

	result.Value = lo.Reduce(mergedRanges, func(agg int, r *Range, index int) int {
		return agg + r.To - r.From + 1
	}, 0)

	return
}

func parse(data string) (ingredients Ingredients) {
	lines := tasks.LineBlocks(data)

	ingredients.FreshRanges = lo.Map(lines[0], func(line string, index int) *Range {
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		return &Range{start, end}
	})
	ingredients.Ingredients = lo.Map(lines[1], func(line string, index int) int {
		id, _ := strconv.Atoi(line)
		return id
	})
	return
}

func mergeRanges(ranges []*Range) []*Range {
	if len(ranges) <= 1 {
		return ranges
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].From < ranges[j].From
	})

	currentRange := ranges[0]
	merged := []*Range{currentRange}

	for _, r := range ranges {
		currentEnd := currentRange.To
		nextStart := r.From
		nextEnd := r.To

		if currentEnd >= nextStart {
			currentRange.To = max(currentEnd, nextEnd)
		} else {
			currentRange = r
			merged = append(merged, currentRange)
		}
	}

	return merged
}
