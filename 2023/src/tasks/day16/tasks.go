package day16

import (
	"2023/src/framework"
	"2023/src/framework/geometry"
	"2023/src/tasks/day16/model"
	"go/types"
	"sync"
)

func Task1(data string, _ types.Nil) (result framework.Result[int]) {
	cave := model.NewCave(data)

	result.Value = cave.CountEnergized(geometry.Point{}, geometry.East)

	return
}

func Task2(data string, _ types.Nil) (result framework.Result[int]) {
	cave := model.NewCave(data)

	var mu sync.RWMutex

	xMin, xMax, yMin, yMax := cave.Boundaries()

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
	wg.Add(xMax - xMin + 1)
	wg.Add(yMax - yMin + 1)

	for x := xMin; x <= xMax; x++ {
		go func(x int) {
			size := cave.CountEnergized(geometry.Point{X: x}, geometry.South)
			setMaxResult(size)

			size = cave.CountEnergized(geometry.Point{X: x, Y: yMax}, geometry.North)
			setMaxResult(size)

			wg.Done()
		}(x)
	}
	for y := yMin; y <= yMax; y++ {
		go func(y int) {
			size := cave.CountEnergized(geometry.Point{X: xMax, Y: y}, geometry.West)
			setMaxResult(size)

			size = cave.CountEnergized(geometry.Point{Y: y}, geometry.East)
			setMaxResult(size)

			wg.Done()
		}(y)
	}

	wg.Wait()

	return
}
