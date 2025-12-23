//go:generate go run aoc/cmd/generate-tests

package day6

import (
	"aoc/framework"
	_math "aoc/framework/math"
	"aoc/framework/tasks"
	"go/types"
	"math"
	"regexp"
	"strings"

	"github.com/samber/lo"
)

// Task1 type:mock 	file:data	expected:288
// Task1 type:real 	file:day6 	expected:252000
func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	lines := tasks.Lines(data)
	races := parseRaces(lines, false)

	result.Value = lo.Reduce(races, func(result int, race Race, index int) int {
		times := make([]int, race.time+1)
		for time := range framework.RangeGen(0, race.time+1, 1) {
			times[time] = time
		}
		return result * lo.CountBy(times, func(time int) bool {
			return time*(race.time-time) > race.distance
		})
	}, 1)

	return
}

// Task2 type:mock 	file:data	expected:71503
// Task2 type:real 	file:day6 	expected:36992486
func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	lines := tasks.Lines(data)
	races := parseRaces(lines, true)
	race := races[0]

	a := float64(-1)
	b := float64(race.time)
	c := float64(race.distance)

	lower := (-b + math.Sqrt(math.Pow(b, 2.0)-(-4*a*c))) / (2 * a)
	upper := (-b - math.Sqrt(math.Pow(b, 2.0)-(-4*a*c))) / (2 * a)

	result.Value = int(math.Floor(upper)-math.Ceil(lower)) + 1

	return
}

var linePattern = regexp.MustCompile(`.+:\s+(.*)`)

func parseRaces(data []string, combine bool) []Race {
	timeMatch := linePattern.FindStringSubmatch(data[0])
	distanceMatch := linePattern.FindStringSubmatch(data[1])

	var times, distances []int

	if combine {
		times = _math.ExtractNumbers(strings.ReplaceAll(timeMatch[1], " ", ""))
		distances = _math.ExtractNumbers(strings.ReplaceAll(distanceMatch[1], " ", ""))
	} else {
		times = _math.ExtractNumbers(timeMatch[1])
		distances = _math.ExtractNumbers(distanceMatch[1])
	}

	return lo.Map(lo.Zip2(times, distances), func(item lo.Tuple2[int, int], index int) Race {
		return Race{item.A, item.B}
	})
}

type Race struct {
	time, distance int
}
