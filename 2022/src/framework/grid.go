package framework

import "fmt"

func DrawPointGrid[V comparable](data map[Point]V, mapping map[V]rune, fallback rune) {
	initalized := false
	var lowestX, lowestY, highestX, highestY int
	for position := range data {
		if !initalized || position.X < lowestX {
			lowestX = position.X
		}
		if !initalized || position.X > highestX {
			highestX = position.X
		}
		if !initalized || position.Y < lowestY {
			lowestY = position.Y
		}
		if !initalized || position.Y > highestY {
			highestY = position.Y
		}
		initalized = true
	}

	for _, y := range Range(lowestY, highestY-lowestY+1, 1) {
		for _, x := range Range(lowestX, highestX-lowestX+1, 1) {
			if value, positionExists := data[Point{X: x, Y: y}]; positionExists {
				if character, valueExists := mapping[value]; valueExists {
					fmt.Print(string(character))
				} else {
					fmt.Print(string(fallback))
				}
			} else {
				fmt.Print(string(fallback))
			}
		}
		fmt.Print("\n")
	}
}
