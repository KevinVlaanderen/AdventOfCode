package day2

import (
	"2023/src/framework"
	"errors"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Task1(filePath string) (result framework.Result[int]) {
	for line := range framework.ReadLines(filePath) {
		game := parse(line)
		if game.valid() {
			result.Value += game.id
		}
	}
	return
}

func Task2(filePath string) (result framework.Result[int]) {
	for line := range framework.ReadLines(filePath) {
		game := parse(line)
		var minRed, minGreen, minBlue int
		for _, hand := range game.hands {
			if hand.red > minRed {
				minRed = hand.red
			}
			if hand.green > minGreen {
				minGreen = hand.green
			}
			if hand.blue > minBlue {
				minBlue = hand.blue
			}
		}
		result.Value += Hand{minRed, minGreen, minBlue}.power()
	}
	return
}

const maxRed, maxGreen, maxBlue = 12, 13, 14

var gamePattern = regexp.MustCompile(`Game (\d+): (.*)`)
var handPattern = regexp.MustCompile(`(\d+) (red|green|blue)`)

type Game struct {
	id    int
	hands []Hand
}

func (game Game) valid() bool {
	for _, hand := range game.hands {
		if hand.red > maxRed || hand.green > maxGreen || hand.blue > maxBlue {
			return false
		}
	}
	return true
}

type Hand struct {
	red, green, blue int
}

func (hand Hand) power() int {
	return hand.red * hand.green * hand.blue
}

func parse(line string) (game Game) {
	gameMatch := gamePattern.FindStringSubmatch(line)
	var err error

	if game.id, err = strconv.Atoi(gameMatch[1]); err != nil {
		panic(err)
	}

	handStrings := strings.Split(gameMatch[2], ";")
	if len(handStrings) == 0 {
		panic(errors.New("no hands found"))
	}

	game.hands = lo.FlatMap(handStrings, func(handString string, index int) []Hand {
		handMatches := handPattern.FindAllStringSubmatch(strings.TrimSpace(handString), -1)

		return lop.Map(handMatches, func(handMatch []string, index int) Hand {
			var red, green, blue, n int
			n, err = strconv.Atoi(handMatch[1])
			if err != nil {
				panic(err)
			}

			switch handMatch[2] {
			case "red":
				red = n
			case "green":
				green = n
			case "blue":
				blue = n
			default:
				log.Panicf("unknown color %v", handMatch[2])
			}

			return Hand{red, green, blue}
		})
	})

	return
}
