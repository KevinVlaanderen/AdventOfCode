package day14

import (
	"2023/src/framework"
	"github.com/samber/lo"
)

func Task1(data string) (result framework.Result[int]) {
	rocks := framework.CharLines(data)

	moveNorth(&rocks)
	result.Value = calculateLoad(&rocks)

	return
}

const MaxIterations = 1000000000

func Task2(data string) (result framework.Result[int]) {
	rocks := framework.CharLines(data)

	var previousHashes = make([]framework.Hash64, 0)
	var previousScores = make([]int, 0)

	for i := 0; i < MaxIterations; i++ {
		moveNorth(&rocks)
		moveWest(&rocks)
		moveSouth(&rocks)
		moveEast(&rocks)

		hash := framework.ComputeStringSliceHash64(lo.Map(rocks, func(item []rune, index int) string {
			return string(item)
		}))

		startOfPeriod := lo.IndexOf(previousHashes, hash)
		if startOfPeriod >= 0 {
			sizeOfPeriod := i - startOfPeriod
			target := (MaxIterations % sizeOfPeriod) - 1

			if target > startOfPeriod {
				result.Value = previousScores[target]
			} else {
				result.Value = previousScores[target+sizeOfPeriod]
			}

			break
		}

		score := calculateLoad(&rocks)

		previousHashes = append(previousHashes, hash)
		previousScores = append(previousScores, score)
	}

	return
}

func calculateLoad(data *[][]rune) (result int) {
	height := len(*data)

	for y, line := range *data {
		for _, char := range line {
			if char == 'O' {
				result += height - y
			}
		}
	}
	return
}

func moveNorth(data *[][]rune) {
	width := len((*data)[0])
	height := len(*data)
	limits := make([]int, width)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch (*data)[y][x] {
			case 'O':
				(*data)[y][x] = '.'
				(*data)[limits[x]][x] = 'O'
				limits[x] += 1
			case '#':
				limits[x] = y + 1
			}
		}
	}
}

func moveWest(data *[][]rune) {
	width := len((*data)[0])
	height := len(*data)
	limits := make([]int, height)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch (*data)[y][x] {
			case 'O':
				(*data)[y][x] = '.'
				(*data)[y][limits[y]] = 'O'
				limits[y] += 1
			case '#':
				limits[y] = x + 1
			}
		}
	}
}

func moveSouth(data *[][]rune) {
	width := len((*data)[0])
	height := len(*data)
	limits := framework.Range(height-1, width, 0)

	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			switch (*data)[y][x] {
			case 'O':
				(*data)[y][x] = '.'
				(*data)[limits[x]][x] = 'O'
				limits[x] -= 1
			case '#':
				limits[x] = y - 1
			}
		}
	}
}

func moveEast(data *[][]rune) {
	width := len((*data)[0])
	height := len(*data)
	limits := framework.Range(width-1, height, 0)

	for x := width - 1; x >= 0; x-- {
		for y := 0; y < height; y++ {
			switch (*data)[y][x] {
			case 'O':
				(*data)[y][x] = '.'
				(*data)[y][limits[y]] = 'O'
				limits[y] -= 1
			case '#':
				limits[y] = x - 1
			}
		}
	}
}
