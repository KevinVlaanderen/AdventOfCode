package day11

import (
	"aoc/2022/tasks/day11/model"
	"aoc/framework"
	"aoc/framework/tasks"
	"go/types"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	monkeys := parse(data)

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

var monkeyPattern = regexp.MustCompile(
	`^Monkey (\d+):\n` +
		`\s+Starting items: (.*)\n` +
		`\s+Operation: new = (.*)\n` +
		`\s+Test: divisible by (\d+)\n` +
		`\s+If true: throw to monkey (\d+)\n` +
		`\s+If false: throw to monkey (\d+)`)

func parse(data string) []model.Monkey {
	return lo.Map(tasks.Blocks(data), func(block string, index int) model.Monkey {
		matches := monkeyPattern.FindStringSubmatch(block)

		number, _ := strconv.Atoi(matches[1])

		var items []model.Item
		itemsString := strings.Split(matches[2], ", ")
		for _, itemString := range itemsString {
			worryLevel, _ := strconv.Atoi(itemString)
			items = append(items, model.Item{WorryLevel: worryLevel})
		}

		operation := model.Expression(matches[3])
		testNumber, _ := strconv.Atoi(matches[4])
		trueTarget, _ := strconv.Atoi(matches[5])
		falseTarget, _ := strconv.Atoi(matches[6])

		return model.Monkey{
			Number:    number,
			Items:     items,
			Operation: operation,
			Test:      testNumber,
			Targets:   map[bool]int{true: trueTarget, false: falseTarget},
		}
	})
}
