package day10

import (
	"aoc/framework/math"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

var machinePattern = regexp.MustCompile(`\[(.+)] (.*) \{(.*)}`)
var buttonWiringSchematicsPattern = regexp.MustCompile(`\d+`)

func Parse(data string) []MachineDescription {
	matches := machinePattern.FindAllStringSubmatch(data, -1)

	return lo.Map(matches, func(match []string, index int) MachineDescription {
		indicatorLightDiagram := lo.Reduce([]rune(match[1]), func(agg int, light rune, index int) int {
			if light == '#' {
				return agg + math.PowInt(2, index)
			}
			return agg
		}, 0)

		buttonWiringSchematics := lo.Map(strings.Split(match[2], " "), func(schematic string, index int) int {
			button := 0
			for _, wiring := range buttonWiringSchematicsPattern.FindAllString(schematic, -1) {
				value, _ := strconv.Atoi(wiring)
				button += math.PowInt(2, value)
			}
			return button
		})

		joltageRequirements := lo.Map(strings.Split(match[3], ","), func(joltage string, index int) int {
			value, _ := strconv.Atoi(joltage)
			return value
		})

		return MachineDescription{
			targetLights:   indicatorLightDiagram,
			buttons:        buttonWiringSchematics,
			targetJoltages: joltageRequirements,
		}
	})
}
