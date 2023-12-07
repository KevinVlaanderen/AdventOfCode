package day7

import (
	"2023/src/framework/task"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

func Task1(filePath string) (result task.Result[int]) {
	hands := task.Read(filePath, parse)

	slices.SortFunc(hands, func(a, b Hand) int {
		return a.Compare(b)
	})

	result.Value = lo.Reduce(hands, func(agg int, hand Hand, index int) int {
		return agg + (index+1)*hand.bid
	}, 0)

	return
}

func parse(line string) (hand Hand, hasResult bool, err error) {
	if line == "" {
		return
	}

	parts := strings.Split(line, " ")
	hand.cards = lo.Map(strings.Split(parts[0], ""), func(item string, index int) Card {
		return Card(item)
	})
	hand.bid, err = strconv.Atoi(parts[1])
	hasResult = true
	return
}

type Hand struct {
	cards []Card
	bid   int
}

func (h Hand) Type() HandType {
	cardCount := lo.CountValues(h.cards)

	switch {
	case len(cardCount) == 1:
		return FiveOfAKind
	case len(cardCount) == 2 && lo.Contains(lo.Values(cardCount), 4):
		return FourOfAKind
	case len(cardCount) == 2 && lo.Contains(lo.Values(cardCount), 3):
		return FullHouse
	case len(cardCount) == 3 && lo.Contains(lo.Values(cardCount), 3):
		return ThreeOfAKind
	case len(cardCount) == 3 && lo.Contains(lo.Values(cardCount), 2):
		return TwoPair
	case len(cardCount) == 4:
		return OnePair
	default:
		return HighCard
	}
}

func (h Hand) Compare(other Hand) int {
	thisType, otherType := h.Type(), other.Type()
	if thisType < otherType {
		return -1
	} else if thisType > otherType {
		return 1
	}
	for _, cards := range lo.Zip2(h.cards, other.cards) {
		aValue, bValue := cards.A.Value(), cards.B.Value()
		if aValue < bValue {
			return -1
		} else if aValue > bValue {
			return 1
		}
	}
	return 0
}

type Card string

func (c Card) Value() int {
	return cardMap[string(c)]
}

var cardMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)
