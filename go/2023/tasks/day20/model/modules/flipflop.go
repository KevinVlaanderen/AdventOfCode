package modules

import "aoc/2023/tasks/day20/model"

type FlipFlop struct {
	*DefaultModule
}

func NewFlipFlop(name string, system *model.System) *FlipFlop {
	return &FlipFlop{DefaultModule: NewDefaultModule(name, system)}
}

func (f *FlipFlop) Receive(source model.Module, value bool) bool {
	f.DefaultModule.Receive(source, value)

	if value {
		return false
	}
	f.value = !f.value
	return true
}
