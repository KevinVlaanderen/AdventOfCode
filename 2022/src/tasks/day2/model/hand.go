package model

type Hand int

const (
	Rock Hand = iota
	Paper
	Scissors
	Invalid
)

var player1HandMap = map[string]Hand{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var player2HandMap = map[string]Hand{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

func ParseHandFor(side PlayerSide, symbol string) Hand {
	var mapping map[string]Hand
	if side == Player1 {
		mapping = player1HandMap
	} else {
		mapping = player2HandMap
	}

	if hand, ok := mapping[symbol]; ok {
		return hand
	} else {
		return Invalid
	}
}

func (h Hand) WinsAgainst() Hand {
	switch h {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	}
	return Invalid
}

func (h Hand) LosesAgainst() Hand {
	switch h {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	}
	return Invalid
}
