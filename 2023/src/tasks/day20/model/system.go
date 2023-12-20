package model

import "github.com/samber/lo"

type State struct {
	LowCount, HighCount int
	Results             map[string]bool
}

type System struct {
	modules map[string]Module
	state   *State
	queue   []Module
}

func NewSystem() System {
	return System{
		modules: make(map[string]Module),
		queue:   make([]Module, 0, 100),
		state:   &State{Results: make(map[string]bool)},
	}
}

func (s *System) SetModules(modules map[string]Module) {
	s.modules = modules
	s.state.Results = make(map[string]bool, len(lo.Keys(modules)))
}

func (s *System) Module(name string) Module {
	return s.modules[name]
}

func (s *System) State() *State {
	return s.state
}

func (s *System) ActivateModule(module Module) {
	currentModule := module
	s.queue = append(s.queue, module)
	for len(s.queue) > 0 {
		currentModule = s.queue[0]
		s.queue = s.queue[1:]
		output := currentModule.Output()
		for _, sink := range currentModule.Sinks() {
			//valueString := "low"
			//if output {
			//	valueString = "high"
			//}
			//fmt.Printf("%v -%v-> %v\n", currentModule.Name(), valueString, sink.Name())

			if output {
				s.state.HighCount++
			} else {
				s.state.LowCount++
			}
			if sink.Receive(currentModule, output) {
				s.queue = append(s.queue, sink)
			}
		}
	}
}

func (s *System) SetResult(name string, value bool) {
	s.state.Results[name] = value
}

type Module interface {
	Name() string
	Sinks() []Module
	Output() bool
	Receive(source Module, value bool) bool
	RegisterSource(sink Module)
	RegisterSink(sink Module)
}
