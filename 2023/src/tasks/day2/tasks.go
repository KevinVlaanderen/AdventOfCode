package day2

import (
	"2023/src/framework/task"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Task1(filePath string) (result task.Result[int]) {
	for game := range task.ReadStream(filePath, parser) {
		if game.valid() {
			result.Value += game.id
		}
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

func parser(line string) (game Game, hasResult bool, err error) {
	if line == "" {
		return
	}

	gameMatch := gamePattern.FindStringSubmatch(line)

	if game.id, err = strconv.Atoi(gameMatch[1]); err != nil {
		return
	}

	handStrings := strings.Split(gameMatch[2], ";")
	if len(handStrings) == 0 {
		err = errors.New("no hands found")
		return
	}

	for _, handString := range handStrings {
		handMatches := handPattern.FindAllStringSubmatch(strings.TrimSpace(handString), -1)

		for _, handMatch := range handMatches {
			var red, green, blue, n int
			n, err = strconv.Atoi(handMatch[1])
			if err != nil {
				return
			}

			switch handMatch[2] {
			case "red":
				red = n
			case "green":
				green = n
			case "blue":
				blue = n
			default:
				err = fmt.Errorf("unknown color %v", handMatch[2])
				return
			}

			game.hands = append(game.hands, Hand{red, green, blue})
		}
	}

	hasResult = true
	return
}
