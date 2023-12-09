package day6

import (
	"2023/src/framework"
	"2023/src/framework/generators"
	"2023/src/framework/number"
	"2023/src/framework/parse"
	"github.com/samber/lo"
	"regexp"
	"strings"
)

func Task1(filePath string) (result framework.Result[int]) {
	data := framework.Read(filePath, parse.Blocks())[0]
	races := parseRaces(data, false)

	result.Value = lo.Reduce(races, func(result int, race Race, index int) int {
		times := make([]int, race.time+1)
		for time := range generators.RangeGen(0, race.time+1, 1) {
			times[time] = time
		}
		return result * lo.CountBy(times, func(time int) bool {
			return time*(race.time-time) > race.distance
		})
	}, 1)

	return
}

func Task2(filePath string) (result framework.Result[int]) {
	data := framework.Read(filePath, parse.Blocks())[0]
	races := parseRaces(data, true)

	result.Value = lo.Reduce(races, func(result int, race Race, index int) int {
		return result * lo.CountBy(generators.Range(0, race.time+1, 1), func(time int) bool {
			return time*(race.time-time) > race.distance
		})
	}, 1)

	return
}

var linePattern = regexp.MustCompile(`.+:\s+(.*)`)

func parseRaces(data []string, combine bool) []Race {
	timeMatch := linePattern.FindStringSubmatch(data[0])
	distanceMatch := linePattern.FindStringSubmatch(data[1])

	var times, distances []int

	if combine {
		times = number.ExtractNumbers(strings.ReplaceAll(timeMatch[1], " ", ""))
		distances = number.ExtractNumbers(strings.ReplaceAll(distanceMatch[1], " ", ""))
	} else {
		times = number.ExtractNumbers(timeMatch[1])
		distances = number.ExtractNumbers(distanceMatch[1])
	}

	return lo.Map(lo.Zip2(times, distances), func(item lo.Tuple2[int, int], index int) Race {
		return Race{item.A, item.B}
	})
}

type Race struct {
	time, distance int
}
