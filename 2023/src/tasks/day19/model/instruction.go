package model

func ParseInstruction(data string) Instruction {
	switch data {
	case "A":
		return Accept()
	case "R":
		return Reject()
	default:
		return Goto(data)
	}
}

type Instruction func(state *State)

func Goto(workflow string) Instruction {
	return func(state *State) {
		state.Workflow = workflow
	}
}

func Accept() Instruction {
	return func(state *State) {
		state.Accepted = true
		state.Done = true
	}
}

func Reject() Instruction {
	return func(state *State) {
		state.Accepted = false
		state.Done = true
	}
}
