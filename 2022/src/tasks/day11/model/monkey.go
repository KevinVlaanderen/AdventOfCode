package model

import (
	"math"
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
