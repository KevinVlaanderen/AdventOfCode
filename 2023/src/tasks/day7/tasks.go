package day7

import (
	"2023/src/framework"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

func Task1(data string) (result framework.Result[int]) {
	hands := parseHands(data, &cardMap1)

	slices.SortFunc(hands, func(a, b Hand) int {
		return a.Compare(b, false)
	})

	result.Value = lo.Reduce(hands, func(agg int, hand Hand, index int) int {
		return agg + (index+1)*hand.bid
	}, 0)

	return
}

func Task2(data string) (result framework.Result[int]) {
	hands := parseHands(data, &cardMap2)

	slices.SortFunc(hands, func(a, b Hand) int {
		return a.Compare(b, true)
	})

	result.Value = lo.Reduce(hands, func(agg int, hand Hand, index int) int {
		return agg + (index+1)*hand.bid
	}, 0)

	return
}

func parseHands(data string, cardMap *map[string]int) []Hand {
	return lo.Map(framework.Lines(data), func(item string, index int) Hand {
		if hand, err := NewHand(item, cardMap); err != nil {
			panic(err)
		} else {
			return hand
		}
	})
}

type Hand struct {
	cardMap *map[string]int
	cards   []Card
	bid     int
}

func NewHand(data string, cardMap *map[string]int) (hand Hand, err error) {
	parts := strings.Split(data, " ")
	hand.cards = lo.Map(strings.Split(parts[0], ""), func(item string, index int) Card {
		return Card(item)
	})
	hand.bid, err = strconv.Atoi(parts[1])
	hand.cardMap = cardMap
	return
}

func (h Hand) Type(withJokers bool) HandType {
	cardCount := lo.CountValues(h.cards)

	jokers := cardCount["J"]
	if withJokers && jokers > 0 {
		withoutJ := lo.Filter(lo.Entries(cardCount), func(item lo.Entry[Card, int], index int) bool {
			return item.Key != "J"
		})
		most := lo.MaxBy(withoutJ, func(entry lo.Entry[Card, int], max lo.Entry[Card, int]) bool {
			return entry.Key != "J" && entry.Value >= max.Value
		}).Key
		delete(cardCount, "J")
		cardCount[most] += jokers
	}

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

func (h Hand) Compare(other Hand, withJokers bool) int {
	thisType, otherType := h.Type(withJokers), other.Type(withJokers)
	if thisType < otherType {
		return -1
	} else if thisType > otherType {
		return 1
	}
	for _, cards := range lo.Zip2(h.cards, other.cards) {
		aValue, bValue := cards.A.Value(h.cardMap), cards.B.Value(h.cardMap)
		if aValue < bValue {
			return -1
		} else if aValue > bValue {
			return 1
		}
	}
	return 0
}

type Card string

func (c Card) Value(cardMap *map[string]int) int {
	return (*cardMap)[string(c)]
}

var cardMap1 = map[string]int{
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

var cardMap2 = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
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
