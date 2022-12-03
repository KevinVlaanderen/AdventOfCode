package model

import "fmt"

type Result int

const (
	Draw Result = iota
	Win
	Loss
)

type Round struct {
	player1 Player
	player2 Player
}

func NewRound(player1Hand Hand, player2Hand Hand) Round {
	return Round{
		player1: NewPlayer(Player1, player1Hand),
		player2: NewPlayer(Player2, player2Hand),
	}
}

func (g Round) ValidHands() bool {
	return g.player1.hand != Invalid && g.player2.hand != Invalid
}

func (g Round) ScoreFor(side PlayerSide) (score int, err error) {
	if !g.ValidHands() {
		return 0, fmt.Errorf("round contains invalid hand")
	}

	player := g.getPlayer(side)

	switch g.resultFor(player) {
	case Loss:
	case Draw:
		score += 3
	case Win:
		score += 6
	}

	switch player.hand {
	case Rock:
		score += 1
	case Paper:
		score += 2
	case Scissors:
		score += 3
	}

	return
}

func (g Round) getPlayer(side PlayerSide) Player {
	if side == Player1 {
		return g.player1
	} else {
		return g.player2
	}
}

func (g Round) getOpponentFor(player Player) Player {
	if player.side == Player1 {
		return g.player2
	} else {
		return g.player1
	}
}

func (g Round) resultFor(player Player) Result {
	opponent := g.getOpponentFor(player)

	switch {
	case player.hand == Rock && opponent.hand == Paper:
		return Loss
	case player.hand == Rock && opponent.hand == Scissors:
		return Win
	case player.hand == Paper && opponent.hand == Rock:
		return Win
	case player.hand == Paper && opponent.hand == Scissors:
		return Loss
	case player.hand == Scissors && opponent.hand == Rock:
		return Loss
	case player.hand == Scissors && opponent.hand == Paper:
		return Win
	}

	return Draw
}

func (g Round) handOf(player PlayerSide) (hand Hand) {
	if player == Player1 {
		hand = g.player1.hand
	} else {
		hand = g.player2.hand
	}
	return
}
