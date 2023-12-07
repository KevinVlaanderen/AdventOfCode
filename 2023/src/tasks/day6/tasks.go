package day6

import (
	"2023/src/framework/generators"
	"2023/src/framework/number"
	"2023/src/framework/parse"
	"2023/src/framework/task"
	"github.com/samber/lo"
	"regexp"
)

func Task1(filePath string) (result task.Result[int]) {
	data := task.Read(filePath, parse.Blocks())[0]
	races := parseRaces(data)

	result.Value = lo.Reduce(races, func(result int, race Race, index int) int {
		return result * lo.CountBy(generators.Range(0, race.time+1, 1), func(time int) bool {
			return time*(race.time-time) > race.distance
		})
	}, 1)

	return
}

var linePattern = regexp.MustCompile(`.+:\s+(.*)`)

func parseRaces(data []string) []Race {
	timeMatch := linePattern.FindStringSubmatch(data[0])
	distanceMatch := linePattern.FindStringSubmatch(data[1])

	times := number.ExtractNumbers(timeMatch[1])
	distances := number.ExtractNumbers(distanceMatch[1])

	return lo.Map(lo.Zip2(times, distances), func(item lo.Tuple2[int, int], index int) Race {
		return Race{item.A, item.B}
	})
}

type Race struct {
	time, distance int
}
