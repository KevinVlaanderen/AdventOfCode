package day11

import (
	"aoc/framework/tasks"
	"go/types"
	"regexp"

	"github.com/oleiade/lane/v2"
	"github.com/samber/lo"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	connections, start, end := Parse(data)
	
	queue := lane.NewQueue[int](start)

	for queue.Size() > 0 {
		current, _ := queue.Dequeue()
		toList := connections[current]

		for _, to := range toList {
			if to == end {
				result.Value++
			} else {
				queue.Enqueue(to)
			}
		}
	}

	return
}

var devicePattern = regexp.MustCompile(`\w+`)

func Parse(data string) (map[int][]int, int, int) {
	lines := tasks.Lines(data)

	deviceIndices := make(map[string]int)
	connections := make(map[int][]int)

	index := 0

	for _, line := range lines {
		devices := devicePattern.FindAllString(line, -1)

		for _, device := range devices {
			if _, ok := deviceIndices[device]; !ok {
				deviceIndices[device] = index
				index++
			}
		}

		fromDevice := devices[0]
		toDevices := devices[1:]

		connections[deviceIndices[fromDevice]] = lo.Map(toDevices, func(device string, index int) int {
			i, _ := deviceIndices[device]
			return i
		})
	}

	start, _ := deviceIndices["you"]
	end, _ := deviceIndices["out"]

	return connections, start, end
}
