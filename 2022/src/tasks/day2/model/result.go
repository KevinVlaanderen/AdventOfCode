package model

type Result = int

const (
	InvalidResult Result = iota
	Draw
	Win
	Loss
)

var resultMap = map[string]Result{
	"X": Loss,
	"Y": Draw,
	"Z": Win,
}

func ParseDesiredResult(symbol string) Result {
	if c, ok := resultMap[symbol]; ok {
		return c
	} else {
		return InvalidResult
	}
}
