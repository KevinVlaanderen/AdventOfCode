package day11

import (
	"2022/src/framework"
	"2022/src/framework/test"
	"2022/src/tasks/day11/model"
	"sort"
)

func Task1(filePath string) (result test.TaskResult[int]) {
	data, err := framework.ReadBlocks(filePath)
	if err != nil {
		result.Error = err
		return
	}

	var monkeys []model.Monkey
	for _, block := range data {
		monkey := model.ParseMonkey(block)
		monkeys = append(monkeys, monkey)
	}

	for range framework.Range(0, 20, 1) {
		for monkeyIndex := range monkeys {
			monkey := &monkeys[monkeyIndex]
			for itemIndex := range monkey.Items {
				item := &monkey.Items[itemIndex]
				target := monkey.Inspect(item)
				targetMonkey := &monkeys[target]
				monkey.Throw(item, targetMonkey)
			}
			monkey.Items = []model.Item{}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspected > monkeys[j].Inspected
	})

	result.Value = monkeys[0].Inspected * monkeys[1].Inspected

	return
}

func Task2(filePath string) (result test.TaskResult[int]) {
	_, err := framework.ReadLineBlocks(filePath)
	if err != nil {
		result.Error = err
		return
	}
	return
}
