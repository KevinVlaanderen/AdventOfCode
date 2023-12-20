package model

import (
	"github.com/oleiade/lane/v2"
	"github.com/samber/lo"
)

type State struct{ LowCount, HighCount int }

type System struct {
	modules           map[string]Module
	registeredSignals *lane.Queue[lo.Tuple3[Module, Module, bool]]
	state             State
}

func NewSystem() System {
	return System{registeredSignals: lane.NewQueue[lo.Tuple3[Module, Module, bool]](), modules: make(map[string]Module)}
}

func (s *System) AddModules(modules map[string]Module) {
	for name, module := range modules {
		s.modules[name] = module
	}
}

func (s *System) Module(name string) Module {
	return s.modules[name]
}

func (s *System) State() State {
	return s.state
}

func (s *System) Tick() {
	for s.registeredSignals.Size() > 0 {
		registeredSignal, _ := s.registeredSignals.Dequeue()
		doSend := registeredSignal.B.Receive(registeredSignal.A, registeredSignal.C)

		//valueString := "low"
		//if registeredSignal.C {
		//	valueString = "high"
		//}
		//fmt.Printf("%v -%v-> %v (new value = %v)\n", registeredSignal.A.Name(), valueString, registeredSignal.B.Name(), registeredSignal.B.Value())

		if doSend {
			registeredSignal.B.Send()
		}
	}
}

func (s *System) RegisterSignal(from Module, to Module, value bool) {
	s.registeredSignals.Enqueue(lo.Tuple3[Module, Module, bool]{A: from, B: to, C: value})
	if value {
		s.state.HighCount++
	} else {
		s.state.LowCount++
	}
}

type Module interface {
	Name() string
	Value() bool
	Send()
	Receive(source Module, value bool) bool
	RegisterSource(sink Module)
	RegisterSink(sink Module)
}
