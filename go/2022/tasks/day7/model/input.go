package model

type InputType int

const (
	COMMAND InputType = iota
	CONTENT
)

type Input struct {
	InputType
	Data string
}

