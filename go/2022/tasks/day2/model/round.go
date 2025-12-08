package model

import "fmt"

type Round struct {
	Player1 Player
	Player2 Player
}

func NewRound(player1Hand Hand, player2Hand Hand) Round {
	return Round{
		Player1: NewPlayer(Player1, player1Hand),
		Player2: NewPlayer(Player2, player2Hand),
	}
}

func (g Round) ValidHands() bool {
	return g.Player1.Hand != Invalid && g.Player2.Hand != Invalid
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

	switch player.Hand {
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
		return g.Player1
	} else {
		return g.Player2
	}
}

func (g Round) getOpponentFor(player Player) Player {
	if player.side == Player1 {
		return g.Player2
	} else {
		return g.Player1
	}
}

func (g Round) resultFor(player Player) Result {
	opponent := g.getOpponentFor(player)

	switch {
	case player.Hand == Rock && opponent.Hand == Paper:
		return Loss
	case player.Hand == Rock && opponent.Hand == Scissors:
		return Win
	case player.Hand == Paper && opponent.Hand == Rock:
		return Win
	case player.Hand == Paper && opponent.Hand == Scissors:
		return Loss
	case player.Hand == Scissors && opponent.Hand == Rock:
		return Loss
	case player.Hand == Scissors && opponent.Hand == Paper:
		return Win
	}

	return Draw
}

func (g Round) handOf(player PlayerSide) (hand Hand) {
	if player == Player1 {
		hand = g.Player1.Hand
	} else {
		hand = g.Player2.Hand
	}
	return
}
