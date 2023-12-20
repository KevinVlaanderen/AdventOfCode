package day20

import (
	"2023/src/framework"
	"2023/src/framework/math"
	"2023/src/tasks/day20/model"
	"2023/src/tasks/day20/model/modules"
	"fmt"
	"github.com/deckarep/golang-set/v2"
	"github.com/dominikbraun/graph"
	"github.com/samber/lo"
	"regexp"
	"strings"
)

func Task1(data string) (result framework.Result[int]) {
	descriptions := CreateDescriptions(data)
	system := CreateSystem(descriptions)

	for i := 0; i < 1000; i++ {
		system.ActivateModule(system.Module("button"))
	}

	result.Value = system.State().LowCount * system.State().HighCount

	return
}

func Task2(data string) (result framework.Result[int]) {
	descriptions := CreateDescriptions(data)

	g := graph.New(graph.StringHash, graph.Rooted(), graph.Directed())
	for name := range descriptions {
		_ = g.AddVertex(name)
	}
	for name, description := range descriptions {
		for _, sink := range description.Sinks {
			_ = g.AddEdge(name, sink)
		}
	}

	adjacencyMap, _ := g.AdjacencyMap()
	predecessorMap, _ := g.PredecessorMap()

	first := lo.Keys(adjacencyMap["broadcaster"])
	final, _ := lo.Find(lo.Keys(adjacencyMap), func(item string) bool {
		return len(lo.Keys(adjacencyMap[item])) == 0
	})
	last := lo.Keys(predecessorMap[final])[0]

	neighbour := first[0]
	counts := make([]int, 1)
	//counts := lop.Map(first, func(neighbour string, index int) int {
	group := []string{"broadcaster", neighbour}
	findGroup(neighbour, adjacencyMap, &group)

	fmt.Printf("From %v to %v %v\n", neighbour, last, group)

	groupDescriptions := make(map[string]Description, len(group))
	for _, name := range group {
		groupDescriptions[name] = descriptions[name]
	}

	system := CreateSystem(groupDescriptions)

	count := 0
	for i := 0; i < 1000; i++ {
		if value, found := system.State().Results[last]; found && value {
			fmt.Printf("Found result %v for neighbour %v", count, neighbour)
			break
		}
		system.ActivateModule(system.Module("broadcaster"))
		count++
		if count%10000000 == 0 {
			fmt.Printf("%v: %v iterations\n", neighbour, count)
		}
	}

	counts[0] = count
	//})

	println(counts)

	result.Value = math.LCM(counts[0], counts[1], counts[2:]...)

	//file, _ := os.Create(path.Join(".", "modules.gv"))
	//_ = draw.DOT(g, file)

	//for !system.State().HasResult || system.State().Result {
	//	system.ActivateModule(system.Module("button"))
	//	result.Value++
	//	if result.Value%1000000 == 0 {
	//		fmt.Printf("%v iterations\n", result.Value)
	//	}
	//}

	return
}

func findGroup(module string, adjacencyMap map[string]map[string]graph.Edge[string], found *[]string) {
	neighbours := lo.Keys(adjacencyMap[module])
	newNeighbours := lo.Filter(neighbours, func(neighbour string, index int) bool {
		return !lo.Contains(*found, neighbour)
	})
	for _, neighbour := range newNeighbours {
		*found = append(*found, neighbour)
	}
	lo.ForEach(newNeighbours, func(neighbour string, index int) {
		findGroup(neighbour, adjacencyMap, found)
	})

	return
}

var modulePattern = regexp.MustCompile(`(?m)^([%&])?(.*) -> (.*)$`)

type ModuleType int

const (
	FlipFlop ModuleType = iota
	Conjunction
	Default
)

type Description struct {
	ModuleType ModuleType
	Sources    []string
	Sinks      []string
}

func CreateDescriptions(data string) map[string]Description {
	moduleMatches := modulePattern.FindAllStringSubmatch(data, -1)

	allModules := make(map[string]Description, len(moduleMatches))
	sourceMapping := make(map[string]mapset.Set[string], len(moduleMatches))
	sinkMapping := make(map[string]mapset.Set[string], len(moduleMatches))

	for _, moduleMatch := range moduleMatches {
		name := moduleMatch[2]

		switch moduleMatch[1] {
		case "%":
			allModules[name] = Description{ModuleType: FlipFlop}
		case "&":
			allModules[name] = Description{ModuleType: Conjunction}
		default:
			allModules[name] = Description{ModuleType: Default}
		}

		if _, found := sourceMapping[name]; !found {
			sourceMapping[name] = mapset.NewSet[string]()
		}
		if _, found := sinkMapping[name]; !found {
			sinkMapping[name] = mapset.NewSet[string]()
		}

		for _, sink := range lo.Map(strings.Split(moduleMatch[3], ","), func(item string, index int) string {
			return strings.TrimSpace(item)
		}) {
			if _, found := sinkMapping[name]; !found {
				sinkMapping[name] = mapset.NewSet[string]()
			}
			sinkMapping[name].Add(sink)

			if _, found := sourceMapping[sink]; !found {
				sourceMapping[sink] = mapset.NewSet[string]()
			}
			sourceMapping[sink].Add(name)
		}
	}

	for _, name := range lo.Filter(lo.Keys(sourceMapping), func(item string, index int) bool {
		_, found := sinkMapping[item]
		return !found
	}) {
		allModules[name] = Description{ModuleType: Default}
		sinkMapping[name] = mapset.NewSet[string]()
	}

	allModules["button"] = Description{ModuleType: Default}
	sourceMapping["button"] = mapset.NewSet[string]()
	sinkMapping["button"] = mapset.NewSet[string]("broadcaster")

	for name, module := range allModules {
		module.Sources = sourceMapping[name].ToSlice()
		allModules[name] = module
		module.Sinks = sinkMapping[name].ToSlice()
		allModules[name] = module
	}

	return allModules
}

func CreateSystem(descriptions map[string]Description) *model.System {
	system := model.NewSystem()

	allModules := make(map[string]model.Module, len(descriptions))

	for name, description := range descriptions {
		var module model.Module
		switch description.ModuleType {
		case FlipFlop:
			module = modules.NewFlipFlop(name, &system)
		case Conjunction:
			module = modules.NewConjunction(name, &system)
		default:
			module = modules.NewDefaultModule(name, &system)
		}

		allModules[name] = module
	}

	for name, description := range descriptions {
		for _, source := range description.Sources {
			if sourceModule, found := allModules[source]; found {
				allModules[name].RegisterSource(sourceModule)
			}
		}
		for _, sink := range description.Sinks {
			if sinkModule, found := allModules[sink]; found {
				allModules[name].RegisterSink(sinkModule)
			}
		}
	}

	system.SetModules(allModules)

	return &system
}
