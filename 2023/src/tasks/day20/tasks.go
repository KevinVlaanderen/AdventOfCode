package day20

import (
	"2023/src/framework"
	"2023/src/tasks/day20/model"
	"2023/src/tasks/day20/model/modules"
	"github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
	"regexp"
	"strings"
)

func Task1(data string) (result framework.Result[int]) {
	system := ParseSystem(data)

	for i := 0; i < 1000; i++ {
		system.Module("button").Send()
		system.Tick()
	}

	result.Value = system.State().LowCount * system.State().HighCount

	return
}

var modulePattern = regexp.MustCompile(`(?m)^([%&])?(.*) -> (.*)$`)

func ParseSystem(data string) *model.System {
	moduleMatches := modulePattern.FindAllStringSubmatch(data, -1)

	system := model.NewSystem()

	allModules := make(map[string]model.Module)
	sourceMapping := make(map[string]mapset.Set[string])
	sinkMapping := make(map[string]mapset.Set[string])

	for _, moduleMatch := range moduleMatches {
		var module model.Module
		switch moduleMatch[1] {
		case "%":
			module = modules.NewFlipFlop(moduleMatch[2], &system)
		case "&":
			module = modules.NewConjunction(moduleMatch[2], &system)
		default:
			module = modules.NewDefaultModule(moduleMatch[2], &system)
		}

		allModules[module.Name()] = module

		if _, found := sourceMapping[module.Name()]; !found {
			sourceMapping[module.Name()] = mapset.NewSet[string]()
		}
		if _, found := sinkMapping[module.Name()]; !found {
			sinkMapping[module.Name()] = mapset.NewSet[string]()
		}

		for _, sink := range lo.Map(strings.Split(moduleMatch[3], ","), func(item string, index int) string {
			return strings.TrimSpace(item)
		}) {
			if _, found := sinkMapping[module.Name()]; !found {
				sinkMapping[module.Name()] = mapset.NewSet[string]()
			}
			sinkMapping[module.Name()].Add(sink)

			if _, found := sourceMapping[sink]; !found {
				sourceMapping[sink] = mapset.NewSet[string]()
			}
			sourceMapping[sink].Add(module.Name())
		}
	}

	for _, name := range lo.Filter(lo.Keys(sourceMapping), func(item string, index int) bool {
		_, found := sinkMapping[item]
		return !found
	}) {
		allModules[name] = modules.NewDefaultModule(name, &system)
		sinkMapping[name] = mapset.NewSet[string]()
	}

	allModules["button"] = modules.NewDefaultModule("button", &system)
	sourceMapping["button"] = mapset.NewSet[string]()
	sinkMapping["button"] = mapset.NewSet[string]("broadcaster")

	for name, module := range allModules {
		for source := range sourceMapping[name].Iter() {
			module.RegisterSource(allModules[source])
		}
		for sink := range sinkMapping[name].Iter() {
			module.RegisterSink(allModules[sink])
		}
	}

	system.AddModules(allModules)

	return &system
}
