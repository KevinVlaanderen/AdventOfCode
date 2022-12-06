package day6

import (
	"2022/src/framework"
	"math/bits"
)

func Task1(filePath string) (result framework.TaskResult[int]) {
	data, err := framework.ReadFull(filePath)
	if err != nil {
		result.Error = err
		return
	}

	marker := findMarker(data, 4)

	return framework.TaskResult[int]{Value: marker}
}

func Task2(filePath string) (result framework.TaskResult[int]) {
	data, err := framework.ReadFull(filePath)
	if err != nil {
		result.Error = err
		return
	}

	marker := findMarker(data, 14)

	return framework.TaskResult[int]{Value: marker}
}

func findMarker(data string, length int) int {
	for i := 0; i < len(data)-length; i++ {
		slice := data[i : i+length]

		var unique uint32
		for _, char := range slice {
			unique |= 1 << (int(char) - 97)
		}
		if bits.OnesCount32(unique) == length {
			return i + length
		}
	}
	return -1
}
