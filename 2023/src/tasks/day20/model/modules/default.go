package modules

import "2023/src/tasks/day20/model"

type DefaultModule struct {
	system  *model.System
	name    string
	value   bool
	sources []model.Module
	sinks   []model.Module
}

func NewDefaultModule(name string, system *model.System) *DefaultModule {
	return &DefaultModule{system: system, name: name, sources: make([]model.Module, 0), sinks: make([]model.Module, 0)}
}

func (d *DefaultModule) Name() string {
	return d.name
}

func (d *DefaultModule) Value() bool {
	return d.value
}

func (d *DefaultModule) Output() bool {
	return d.value
}

func (d *DefaultModule) Sinks() []model.Module {
	return d.sinks
}

func (d *DefaultModule) Receive(_ model.Module, value bool) bool {
	d.system.SetResult(d.name, value)
	return true
}

func (d *DefaultModule) RegisterSource(sink model.Module) {
	d.sources = append(d.sources, sink)
}

func (d *DefaultModule) RegisterSink(sink model.Module) {
	d.sinks = append(d.sinks, sink)
}
