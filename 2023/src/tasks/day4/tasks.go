package day4

import (
	"2023/src/framework"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"go/types"
	"golang.org/x/exp/slices"
	"regexp"
	"strconv"
)

func Task1(data string, _ types.Nil) (result framework.Result[int]) {
	cards, err := parse(data)
	if err != nil {
		result.Error = err
		return
	}
	for _, card := range cards {
		result.Value += card.Score()
	}
	return
}

func Task2(data string, _ types.Nil) (result framework.Result[int]) {
	cards, err := parse(data)
	if err != nil {
		result.Error = err
		return
	}

	pile := lo.Map(cards, func(item Card, index int) int {
		return 1
	})
	for index, count := range pile {
		score := cards[index].Correct()
		for nextIndex := range framework.RangeGen(index+1, score, 1) {
			if nextIndex < len(cards) {
				pile[nextIndex] += count
			}
		}
	}

	result.Value = lo.Sum(pile)
	return
}

var cardPattern = regexp.MustCompile(`Card\s+(\d+): ([\d\s]+) \| ([\d\s]+)`)
var numberPattern = regexp.MustCompile(`\d+`)

type Card struct {
	id             int
	winningNumbers []int
	yourNumbers    []int
}

func NewCard(data string) (card Card, err error) {
	cardMatch := cardPattern.FindStringSubmatch(data)

	if card.id, err = strconv.Atoi(cardMatch[1]); err != nil {
		return
	}

	card.winningNumbers = mapStringsToNumbers(numberPattern.FindAllString(cardMatch[2], -1))
	card.yourNumbers = mapStringsToNumbers(numberPattern.FindAllString(cardMatch[3], -1))

	return
}

func (c Card) Score() int {
	correct := c.Correct()
	switch c.Correct() {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return 1 << (correct - 1)
	}
}

func (c Card) Correct() int {
	return lo.CountBy(c.yourNumbers, func(item int) bool {
		return slices.Contains(c.winningNumbers, item)
	})
}

func parse(data string) (cards []Card, err error) {
	for _, line := range framework.Lines(data) {
		var card Card
		if card, err = NewCard(line); err != nil {
			return
		} else {
			cards = append(cards, card)
		}
	}
	return
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
