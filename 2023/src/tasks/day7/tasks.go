package day7

import (
	"2023/src/framework"
	"github.com/samber/lo"
	"go/types"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

func Task1(data string, _ types.Nil) (result framework.Result[int]) {
	hands := parseHands(data, &cardMap1, false)

	slices.SortFunc(hands, func(a, b Hand) int {
		return a.Compare(b)
	})

	result.Value = lo.Reduce(hands, func(agg int, hand Hand, index int) int {
		return agg + (index+1)*hand.bid
	}, 0)

	return
}

func Task2(data string, _ types.Nil) (result framework.Result[int]) {
	hands := parseHands(data, &cardMap2, true)

	slices.SortFunc(hands, func(a, b Hand) int {
		return a.Compare(b)
	})

	result.Value = lo.Reduce(hands, func(agg int, hand Hand, index int) int {
		return agg + (index+1)*hand.bid
	}, 0)

	return
}

func parseHands(data string, cardMap *map[string]int, withJokers bool) []Hand {
	return lo.Map(framework.Lines(data), func(item string, index int) Hand {
		if hand, err := NewHand(item, cardMap, withJokers); err != nil {
			panic(err)
		} else {
			return hand
		}
	})
}

type Hand struct {
	cardMap  *map[string]int
	cards    []Card
	bid      int
	handType HandType
}

func NewHand(data string, cardMap *map[string]int, withJokers bool) (hand Hand, err error) {
	cardsData := data[:5]
	bidData := data[6:]
	hand.cards = lo.Map(strings.Split(cardsData, ""), func(item string, index int) Card {
		return Card(item)
	})
	hand.bid, err = strconv.Atoi(bidData)
	hand.cardMap = cardMap
	hand.handType = calculateType(hand.cards, withJokers)
	return
}

func calculateType(cards []Card, withJokers bool) HandType {
	cardCount := lo.CountValues(cards)

	jokers := cardCount["J"]
	if withJokers && jokers > 0 {
		var most Card
		var mostCount int

		for key, value := range cardCount {
			if key == "J" {
				continue
			}

			if value >= mostCount {
				most = key
				mostCount = value
			}
		}

		delete(cardCount, "J")
		cardCount[most] += jokers
	}

	count := len(cardCount)
	values := lo.Values(cardCount)

	switch {
	case count == 1:
		return FiveOfAKind
	case count == 2 && lo.Contains(values, 4):
		return FourOfAKind
	case count == 2 && lo.Contains(values, 3):
		return FullHouse
	case count == 3 && lo.Contains(values, 3):
		return ThreeOfAKind
	case count == 3 && lo.Contains(values, 2):
		return TwoPair
	case count == 4:
		return OnePair
	default:
		return HighCard
	}
}

func (h Hand) Compare(other Hand) int {
	thisType, otherType := h.handType, other.handType
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
