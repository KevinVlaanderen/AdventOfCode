package model

type PlayerSide int

const (
	Player1 PlayerSide = iota
	Player2
)

type Player struct {
	side PlayerSide
	Hand Hand
}

func NewPlayer(side PlayerSide, hand Hand) (player Player) {
	return Player{side: side, Hand: hand}
}
