package day10

import (
	"aoc/framework/math"
	"aoc/framework/tasks"
	"fmt"
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

	result.Value = int(lo.SumBy(machines, func(machine MachineDescription) uint {
		buttonsPressed := pressButtonsUntilValidLights(machine.buttons, machine.targetLights)
		bestButtonsPressed := lo.MinBy(buttonsPressed, func(a uint, b uint) bool {
			return bits.OnesCount(a) < bits.OnesCount(b)
		})
		return uint(bits.OnesCount(bestButtonsPressed))
	}))

	return
}

func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	machines := Parse(data)

	result.Value = int(lo.SumBy(machines, func(machine MachineDescription) uint {
		buttonsPressed, _ := pressButtonsUntilValidJoltages(machine.buttons, machine.targetJoltages, 0, 0, buildButtonsPressedMaskMap(machine))
		return buttonsPressed
	}))

	return
}

func buildButtonsPressedMaskMap(machine MachineDescription) map[uint][]uint {
	type State struct {
		buttonIndex        int
		lightsMask         uint
		buttonsPressedMask uint
	}

	buttonsPressedMaskMap := make(map[uint][]uint)

	queue := lane.NewQueue[State](State{})

	for queue.Size() > 0 {
		state, _ := queue.Dequeue()

		if state.buttonIndex < len(machine.buttons) {
			queue.Enqueue(State{
				buttonIndex:        state.buttonIndex + 1,
				lightsMask:         state.lightsMask,
				buttonsPressedMask: state.buttonsPressedMask,
			})
			queue.Enqueue(State{
				buttonIndex:        state.buttonIndex + 1,
				lightsMask:         state.lightsMask ^ machine.buttons[state.buttonIndex],
				buttonsPressedMask: state.buttonsPressedMask | uint(math.PowInt(2, state.buttonIndex)),
			})
		} else {
			if _, ok := buttonsPressedMaskMap[state.lightsMask]; ok {
				buttonsPressedMaskMap[state.lightsMask] = append(buttonsPressedMaskMap[state.lightsMask], state.buttonsPressedMask)
			} else {
				buttonsPressedMaskMap[state.lightsMask] = []uint{state.buttonsPressedMask}
			}
		}
	}

	return buttonsPressedMaskMap
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
	answers := make(map[uint]struct{})

	type State struct {
		lightsMask         uint
		buttonsPressedMask uint
		startButtonIndex   uint
	}

	queue := lane.NewQueue(State{})

	for queue.Size() > 0 {
		state, _ := queue.Dequeue()

		if state.lightsMask == target {
			answers[state.buttonsPressedMask] = struct{}{}
			continue
		}

		for index := state.startButtonIndex; index < uint(len(buttons)); index++ {
			targetLights := buttons[index]

			newLightsMark := state.lightsMask ^ targetLights
			buttonMask := uint(math.PowInt(2, int(index)))
			newButtonsPressedMask := state.buttonsPressedMask | buttonMask

			queue.Enqueue(State{lightsMask: newLightsMark, buttonsPressedMask: newButtonsPressedMask, startButtonIndex: index + 1})
		}
	}

	return lo.Keys(answers)
}

type Target []uint

func (t Target) ToString() string {
	return fmt.Sprintf("{%v}", lo.Map(t, func(counter uint, index int) string {
		return strconv.Itoa(int(counter))
	}))
}

func pressButtonsUntilValidJoltages(buttonTargetsMasks []uint, target Target, best uint, level uint, buttonsPressedMaskMap map[uint][]uint) (uint, bool) {
	if lo.EveryBy(target, func(joltage uint) bool {
		return joltage == 0
	}) {
		return 0, true
	}

	if !lo.EveryBy(target, func(joltage uint) bool {
		return joltage <= 1000
	}) {
		return 0, false
	}

	targetLightsMask := lo.Reduce(target, func(agg uint, joltage uint, index int) uint {
		if joltage%2 == 0 {
			return agg
		}
		return agg | uint(math.PowInt(2, index))
	}, 0)

	buttonsPressedMasks, _ := buttonsPressedMaskMap[targetLightsMask]

	// A hack to still find a solution if none of the initial options lead to a solution
	if level == 0 {
		for i := 0; i < len(buttonTargetsMasks); i++ {
			buttonsPressedMasks = append(buttonsPressedMasks, uint(math.PowInt(2, i)))
		}
	}

	found := false

buttonsPressedLoop:
	for _, buttonsPressedMask := range buttonsPressedMasks {
		numButtonsPressed := uint(bits.OnesCount(buttonsPressedMask))

		newTarget := make(Target, len(target))
		copy(newTarget, target)

		for bIndex, buttonTargetsMask := range buttonTargetsMasks {
			if buttonsPressedMask&uint(math.PowInt(2, bIndex)) > 0 {
				for tIndex := range newTarget {
					if buttonTargetsMask&uint(math.PowInt(2, tIndex)) > 0 {
						newTarget[tIndex]--
						if newTarget[tIndex] >= 1000 {
							continue buttonsPressedLoop
						}
					}
				}
			}
		}

		mult := uint(1)

		if !lo.EveryBy(newTarget, func(joltage uint) bool {
			return joltage == 0
		}) {
			for lo.EveryBy(newTarget, func(joltage uint) bool {
				return joltage%2 == 0
			}) {
				newTarget = lo.Map(newTarget, func(joltage uint, index int) uint {
					return joltage / 2
				})
				mult *= 2
			}
		}

		result, ok := pressButtonsUntilValidJoltages(buttonTargetsMasks, newTarget, best, level+1, buttonsPressedMaskMap)
		if ok {
			total := numButtonsPressed + mult*result
			if best == 0 || total < best {
				found = true
				best = total
			}
		}
	}

	if found {
		return best, true
	}

	return 0, false
}
