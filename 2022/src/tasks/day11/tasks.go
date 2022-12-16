package day11

import (
	"2022/src/framework/generators"
	"2022/src/framework/tasks"
	"2022/src/tasks/day11/model"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Task1(filePath string) (result tasks.TaskResult[int]) {
	monkeys := tasks.Read(filePath, createParser())

	for range generators.Range(0, 20, 1) {
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

func Task2(filePath string) (result tasks.TaskResult[int]) {
	return
}

var monkeyPattern = regexp.MustCompile(
	`^Monkey (\d+):\n` +
		`\s+Starting items: (.*)\n` +
		`\s+Operation: new = (.*)\n` +
		`\s+Test: divisible by (\d+)\n` +
		`\s+If true: throw to monkey (\d+)\n` +
		`\s+If false: throw to monkey (\d+)`)

func createParser() func(line string) (result model.Monkey, hasResult bool, err error) {
	var lines []string

	return func(line string) (result model.Monkey, hasResult bool, err error) {
		if line != "" {
			lines = append(lines, line)
			return result, false, nil
		}

		matches := monkeyPattern.FindStringSubmatch(strings.Join(lines, "\n"))

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

		lines = nil

		return model.Monkey{
			Number:    number,
			Items:     items,
			Operation: operation,
			Test:      testNumber,
			Targets:   map[bool]int{true: trueTarget, false: falseTarget},
		}, true, nil
	}
}