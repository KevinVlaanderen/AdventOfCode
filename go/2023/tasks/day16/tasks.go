package day16

import (
	"aoc/2023/tasks/day16/model"
	"aoc/framework/geometry/geo2d"
	"aoc/framework/tasks"
	"go/types"
	"sync"
)

func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	cave := model.NewCave(data)

	result.Value = cave.CountEnergized(geo2d.Point{}, geo2d.East)

	return
}

func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	cave := model.NewCave(data)

	var mu sync.RWMutex

	minX, minY, maxX, maxY := cave.Grid.Bounds()

	getResult := func() int {
		mu.RLock()
		defer mu.RUnlock()
		return result.Value
	}

	setResult := func(value int) {
		mu.Lock()
		defer mu.Unlock()
		result.Value = value
	}

	setMaxResult := func(value int) {
		maxResult := getResult()
		if maxResult == 0 || value > maxResult {
			setResult(value)
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(maxX - minX + 1)
	wg.Add(maxY - minY + 1)

	for x := minX; x <= maxX; x++ {
		go func(x int) {
			size := cave.CountEnergized(geo2d.Point{X: x}, geo2d.South)
			setMaxResult(size)

			size = cave.CountEnergized(geo2d.Point{X: x, Y: maxY}, geo2d.North)
			setMaxResult(size)

			wg.Done()
		}(x)
	}
	for y := minY; y <= maxY; y++ {
		go func(y int) {
			size := cave.CountEnergized(geo2d.Point{X: maxX, Y: y}, geo2d.West)
			setMaxResult(size)

			size = cave.CountEnergized(geo2d.Point{Y: y}, geo2d.East)
			setMaxResult(size)

			wg.Done()
		}(y)
	}

	wg.Wait()

	return
}
