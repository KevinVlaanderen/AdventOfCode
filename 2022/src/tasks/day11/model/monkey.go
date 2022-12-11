package model

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Monkey struct {
	Number    int
	Items     []Item
	Operation Expression
	Test      int
	Targets   map[bool]int

	Inspected int
}

func (m *Monkey) Inspect(item *Item) int {
	m.Inspected++

	context := map[string]int{"old": item.WorryLevel}
	result, _ := m.Operation.Perform(context)
	item.WorryLevel = int(math.Floor(float64(result) / 3))

	testPassed := item.WorryLevel%m.Test == 0

	return m.Targets[testPassed]
}

func (m *Monkey) Throw(item *Item, monkey *Monkey) {
	monkey.Catch(item)
}

func (m *Monkey) Catch(item *Item) {
	m.Items = append(m.Items, *item)
}

var monkeyPattern = regexp.MustCompile(
	`^Monkey (\d+):\n` +
		`\s+Starting items: (.*)\n` +
		`\s+Operation: new = (.*)\n` +
		`\s+Test: divisible by (\d+)\n` +
		`\s+If true: throw to monkey (\d+)\n` +
		`\s+If false: throw to monkey (\d+)`)

func ParseMonkey(data string) Monkey {
	matches := monkeyPattern.FindStringSubmatch(data)

	number, _ := strconv.Atoi(matches[1])

	var items []Item
	itemsString := strings.Split(matches[2], ", ")
	for _, itemString := range itemsString {
		worryLevel, _ := strconv.Atoi(itemString)
		items = append(items, Item{WorryLevel: worryLevel})
	}

	operation := Expression(matches[3])

	test, _ := strconv.Atoi(matches[4])

	trueTarget, _ := strconv.Atoi(matches[5])
	falseTarget, _ := strconv.Atoi(matches[6])

	return Monkey{
		Number:    number,
		Items:     items,
		Operation: operation,
		Test:      test,
		Targets:   map[bool]int{true: trueTarget, false: falseTarget},
	}
}
