package modules

import "2023/src/tasks/day20/model"

type Conjunction struct {
	*DefaultModule
	memValues map[string]bool
}

func NewConjunction(name string, system *model.System) *Conjunction {
	return &Conjunction{DefaultModule: NewDefaultModule(name, system), memValues: make(map[string]bool)}
}

func (c *Conjunction) Receive(source model.Module, value bool) bool {
	c.memValues[source.Name()] = value
	return true
}

func (c *Conjunction) Send() {
	allHigh := true
	for _, module := range c.sources {
		latestValue, found := c.memValues[module.Name()]
		if !found {
			latestValue = module.Value()
		}
		if !latestValue {
			allHigh = false
			break
		}
	}
	for _, sink := range c.sinks {
		c.system.RegisterSignal(c, sink, !allHigh)
	}
}
