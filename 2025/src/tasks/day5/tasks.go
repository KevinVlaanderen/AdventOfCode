package day5

import (
	"2025/src/framework"
	"go/types"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Ingredients struct {
	FreshRanges []Range
	Ingredients []int
}

type Range struct {
	From, To int
}

func Task1(data string, _ types.Nil) (result framework.Result[int]) {
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

func parse(data string) (ingredients Ingredients) {
	lines := framework.LineBlocks(data)

	ingredients.FreshRanges = lo.Map(lines[0], func(line string, index int) Range {
		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		return Range{start, end}
	})
	ingredients.Ingredients = lo.Map(lines[1], func(line string, index int) int {
		id, _ := strconv.Atoi(line)
		return id
	})
	return
}
