package day10

import (
	"aoc/framework/math"
	"aoc/framework/tasks"
	"go/types"
	"math/bits"
	"regexp"
	"strconv"
	"strings"

	"github.com/oleiade/lane/v2"
	"github.com/samber/lo"
)

type MachineDescription struct {
	buttons        []uint
	targetLights   uint
	targetJoltages []uint
}

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	machines := Parse(data)

	result.Value = lo.SumBy(machines, func(machine MachineDescription) int {
		buttonsPressed := pressButtonsUntilValidLights(machine.buttons, machine.targetLights)
		bestButtonsPressed := lo.MinBy(buttonsPressed, func(a uint, b uint) bool {
			return bits.OnesCount(a) < bits.OnesCount(b)
		})
		return bits.OnesCount(bestButtonsPressed)
	})

	return
}

var machinePattern = regexp.MustCompile(`\[(.+)] (.*) \{(.*)}`)
var buttonWiringSchematicsPattern = regexp.MustCompile(`\d+`)

func Parse(data string) []MachineDescription {
	matches := machinePattern.FindAllStringSubmatch(data, -1)

	return lo.Map(matches, func(match []string, index int) MachineDescription {
		buttonWiringSchematics := lo.Map(strings.Split(match[2], " "), func(schematic string, index int) uint {
			button := 0
			for _, wiring := range buttonWiringSchematicsPattern.FindAllString(schematic, -1) {
				value, _ := strconv.Atoi(wiring)
				button += math.PowInt(2, value)
			}
			return uint(button)
		})

		indicatorLightDiagram := lo.Reduce([]rune(match[1]), func(agg uint, light rune, index int) uint {
			if light == '#' {
				return agg + uint(math.PowInt(2, index))
			}
			return agg
		}, 0)

		joltageRequirements := lo.Map(strings.Split(match[3], ","), func(joltage string, index int) uint {
			value, _ := strconv.Atoi(joltage)
			return uint(value)
		})

		return MachineDescription{
			buttons:        buttonWiringSchematics,
			targetLights:   indicatorLightDiagram,
			targetJoltages: joltageRequirements,
		}
	})
}

func pressButtonsUntilValidLights(buttons []uint, target uint) []uint {
	answers := make([]uint, 0)
	buttonPressesSeen := make(map[uint]struct{})

	type State struct {
		lights         uint
		presses        uint
		buttonsPressed uint
	}

	queue := lane.NewQueue(State{lights: 0, presses: 0, buttonsPressed: 0})

	for queue.Size() > 0 {
		state, _ := queue.Dequeue()

		if _, ok := buttonPressesSeen[state.buttonsPressed]; ok {
			continue
		}
		buttonPressesSeen[state.buttonsPressed] = struct{}{}

		if state.lights == target {
			answers = append(answers, state.buttonsPressed)
			continue
		}

		for index, targetLights := range buttons {
			newPresses := state.presses + 1
			newLights := state.lights ^ targetLights
			newButtonsPressed := state.buttonsPressed | uint(math.PowInt(2, index))

			queue.Enqueue(State{lights: newLights, presses: newPresses, buttonsPressed: newButtonsPressed})
		}
	}

	return answers
}
