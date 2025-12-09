package day3

import (
	"aoc/2022/tasks/day3/model"
	"aoc/framework/datastructures"
	"aoc/framework/tasks"
	"fmt"
	"go/types"

	"github.com/samber/lo"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	rucksacks := parse1(data)

	for _, rucksack := range rucksacks {
		if difference, err := datastructures.Intersection(rucksack.Compartment1(), rucksack.Compartment2()); err != nil {
			result.Error = err
			return
		} else {
			if len(difference) != 1 {
				result.Error = fmt.Errorf("invalid number of differences in %s", string(rucksack))
				return
			}

			result.Value += difference[0].GetPriority()
		}
	}

	return
}

func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	groups := parse2(data)

	for _, group := range groups {
		if difference, err := datastructures.Intersection(group[0], group[1], group[2]); err != nil {
			result.Error = err
			return
		} else {
			if len(difference) != 1 {
				result.Error = fmt.Errorf("invalid number of differences in group %v", group)
				return
			}

			result.Value += difference[0].GetPriority()
		}
	}

	return
}

func parse1(data string) []model.Rucksack {
	return lo.Map(tasks.Lines(data), func(line string, index int) model.Rucksack {
		return model.Rucksack(line)
	})
}

func parse2(data string) (groups []model.Group) {
	var group model.Group

	for _, line := range tasks.Lines(data) {
		group = append(group, model.Rucksack(line))
		if len(group) == 3 {
			groups = append(groups, group)
			group = model.Group{}
		}
	}

	return
}
