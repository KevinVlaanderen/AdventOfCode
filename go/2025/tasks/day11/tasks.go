package day11

import (
	"aoc/framework/tasks"
	"go/types"
	"regexp"

	"github.com/oleiade/lane/v2"
	"github.com/samber/lo"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	connections, deviceIndices := Parse(data)

	start, _ := deviceIndices["you"]
	end, _ := deviceIndices["out"]

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

func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	connections, deviceIndices := Parse(data)

	start, _ := deviceIndices["svr"]
	end, _ := deviceIndices["out"]
	dac, _ := deviceIndices["dac"]
	fft, _ := deviceIndices["fft"]

	result.Value = lo.Sum(lo.Map([][]int{{dac, fft}, {fft, dac}}, func(gates []int, index int) int {
		return findPath(connections, start, gates[0], make(map[int]int)) * findPath(connections, gates[0], gates[1], make(map[int]int)) * findPath(connections, gates[1], end, make(map[int]int))
	}))

	return
}

func findPath(connections map[int][]int, from, to int, cache map[int]int) int {
	if result, ok := cache[from]; ok {
		return result
	}

	nextList, _ := connections[from]
	total := 0

	for _, next := range nextList {
		if next == to {
			total += 1
		} else {
			total += findPath(connections, next, to, cache)
		}
	}

	cache[from] = total

	return total
}

var devicePattern = regexp.MustCompile(`\w+`)

func Parse(data string) (map[int][]int, map[string]int) {
	lines := tasks.Lines(data)

	connections := make(map[int][]int)
	deviceIndices := make(map[string]int)

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

	return connections, deviceIndices
}
