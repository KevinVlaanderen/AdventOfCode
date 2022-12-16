package day3

import (
	"2022/src/framework/sets"
	"2022/src/framework/tasks"
	"2022/src/tasks/day3/model"
	"fmt"
)

func Task1(filePath string) (result tasks.TaskResult[int]) {
	rucksacks := tasks.ReadStream(filePath, task1Parser)

	for rucksack := range rucksacks {
		if difference, err := sets.Intersection(rucksack.Compartment1(), rucksack.Compartment2()); err != nil {
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

func Task2(filePath string) (result tasks.TaskResult[int]) {
	groups := tasks.ReadStream(filePath, createTask2Parser())

	for group := range groups {
		if difference, err := sets.Intersection(group[0], group[1], group[2]); err != nil {
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

func task1Parser(line string) (result model.Rucksack, hasResult bool, err error) {
	if line == "" {
		return result, false, nil
	}
	return model.Rucksack(line), true, nil
}

func createTask2Parser() func(line string) (result model.Group, hasResult bool, err error) {
	var group model.Group

	return func(line string) (result model.Group, hasResult bool, err error) {
		if line == "" {
			return result, false, nil
		}

		group = append(group, model.Rucksack(line))
		if len(group) == 3 {
			result = group
			group = nil
			return result, true, nil
		} else {
			return result, false, nil
		}
	}
}

