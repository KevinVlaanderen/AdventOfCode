package day6

import (
	"2023/src/framework"
	"2023/src/framework/math"
	"github.com/samber/lo"
	"regexp"
	"strings"
)

func Task1(data string) (result framework.Result[int]) {
	lines := framework.Lines(data)
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

func Task2(data string) (result framework.Result[int]) {
	lines := framework.Lines(data)
	races := parseRaces(lines, true)
	race := races[0]

	for time := 0; time < race.time; time++ {
		if time*(race.time-time) > race.distance {
			result.Value++
		}
	}

	return
}

var linePattern = regexp.MustCompile(`.+:\s+(.*)`)

func parseRaces(data []string, combine bool) []Race {
	timeMatch := linePattern.FindStringSubmatch(data[0])
	distanceMatch := linePattern.FindStringSubmatch(data[1])

	var times, distances []int

	if combine {
		times = math.ExtractNumbers(strings.ReplaceAll(timeMatch[1], " ", ""))
		distances = math.ExtractNumbers(strings.ReplaceAll(distanceMatch[1], " ", ""))
	} else {
		times = math.ExtractNumbers(timeMatch[1])
		distances = math.ExtractNumbers(distanceMatch[1])
	}

	return lo.Map(lo.Zip2(times, distances), func(item lo.Tuple2[int, int], index int) Race {
		return Race{item.A, item.B}
	})
}

type Race struct {
	time, distance int
}
