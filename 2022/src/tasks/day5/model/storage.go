package model

import "regexp"

var numberPattern = regexp.MustCompile(`\d+`)

type Storage struct {
	Stacks []Stack
}

func NewStorage(data []string) (storage *Storage) {
	numberLine := data[len(data)-1]
	stackData := data[0 : len(data)-1]
	matches := numberPattern.FindAllString(numberLine, -1)
	numStacks := len(matches)

	storage = &Storage{Stacks: make([]Stack, numStacks)}

	for lineIndex := len(stackData) - 1; lineIndex >= 0; lineIndex-- {
		line := stackData[lineIndex]
		for stackIndex := 0; stackIndex < len(storage.Stacks); stackIndex++ {
			crateIndex := stackIndex * 4
			crateString := line[crateIndex : crateIndex+3]

			if crate, err := NewCrate(crateString); err == nil {
				storage.Stacks[stackIndex].Add(crate)
			}
		}
	}

	return
}

func (s *Storage) Move(from int, to int) {
	crate := s.Stacks[from].Remove()
	s.Stacks[to].Add(crate)
}

func (s *Storage) ExecuteInstruction(instruction Instruction) {
	for i := 0; i < instruction.N; i++ {
		s.Move(instruction.From, instruction.To)
	}
}
