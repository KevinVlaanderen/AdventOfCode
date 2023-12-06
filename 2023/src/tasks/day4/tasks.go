package day4

import (
	"2023/src/framework/task"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"golang.org/x/exp/slices"
	"regexp"
	"strconv"
)

func Task1(filePath string) (result task.Result[int]) {
	for card := range task.ReadStream(filePath, parser) {
		result.Value += card.Score()
	}
	return
}

var cardPattern = regexp.MustCompile(`Card\s+(\d+): ([\d\s]+) \| ([\d\s]+)`)
var numberPattern = regexp.MustCompile(`\d+`)

type Card struct {
	id             int
	winningNumbers []int
	yourNumbers    []int
}

func (c Card) Score() int {
	correct := lo.CountBy(c.yourNumbers, func(item int) bool {
		return slices.Contains(c.winningNumbers, item)
	})
	switch correct {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return 1 << (correct - 1)
	}
}

func parser(line string) (card Card, hasResult bool, err error) {
	if line == "" {
		return
	}

	cardMatch := cardPattern.FindStringSubmatch(line)

	if card.id, err = strconv.Atoi(cardMatch[1]); err != nil {
		return
	}

	card.winningNumbers = mapStringsToNumbers(numberPattern.FindAllString(cardMatch[2], -1))
	card.yourNumbers = mapStringsToNumbers(numberPattern.FindAllString(cardMatch[3], -1))

	return card, true, nil
}

func mapStringsToNumbers(strings []string) []int {
	return lop.Map(strings, func(item string, index int) int {
		if number, err := strconv.Atoi(item); err != nil {
			panic(err)
		} else {
			return number
		}
	})
}
