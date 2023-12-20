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

func (d *DefaultModule) RegisterSource(sink model.Module) {
	d.sources = append(d.sources, sink)
}

func (d *DefaultModule) RegisterSink(sink model.Module) {
	d.sinks = append(d.sinks, sink)
}

func (d *DefaultModule) Receive(_ model.Module, _ bool) bool {
	return true
}

func (d *DefaultModule) Send() {
	for _, sink := range d.sinks {
		d.system.RegisterSignal(d, sink, d.value)
	}
}
