package modules

import "2023/src/tasks/day20/model"

type Conjunction struct {
	*DefaultModule
	memValues map[model.Module]bool
}

func NewConjunction(name string, system *model.System) *Conjunction {
	return &Conjunction{DefaultModule: NewDefaultModule(name, system), memValues: make(map[model.Module]bool)}
}

func (c *Conjunction) Receive(source model.Module, value bool) bool {
	c.DefaultModule.Receive(source, value)

	c.memValues[source] = value
	return true
}

func (c *Conjunction) Output() bool {
	for _, module := range c.sources {
		if latestValue, found := c.memValues[module]; !found || !latestValue {
			return true
		}
	}
	return false
}
