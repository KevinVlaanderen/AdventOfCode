package model

func ParseSystem(data string) System {
	workflows := ParseWorkflows(data)
	return System{Workflows: workflows}
}

type System struct {
	Workflows map[string][]Rule
}

func (s *System) Check(part Part) bool {
	state := State{Part: part, Workflow: "in"}

	for !state.Done {
		rules := s.Workflows[state.Workflow]

		for _, rule := range rules {
			if rule.Condition(&state) {
				rule.Instruction(&state)
				break
			}
		}
	}

	return state.Accepted
}

type State struct {
	Part           Part
	Workflow       string
	Done, Accepted bool
}
