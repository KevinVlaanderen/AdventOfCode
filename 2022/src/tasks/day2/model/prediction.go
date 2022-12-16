package model

type Prediction struct {
	Player1 Player
	Outcome Result
}

func NewPrediction(player1Hand Hand, outcome Result) Prediction {
	return Prediction{
		Player1: NewPlayer(Player1, player1Hand),
		Outcome: outcome,
	}
}

func (p Prediction) RequiredHand() Hand {
	switch p.Outcome {
	case Win:
		return p.Player1.Hand.LosesAgainst()
	case Loss:
		return p.Player1.Hand.WinsAgainst()
	case Draw:
		return p.Player1.Hand
	}
	return Invalid
}
