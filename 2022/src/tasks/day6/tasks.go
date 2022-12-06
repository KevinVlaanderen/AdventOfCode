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

	for i := 0; i < len(data)-4; i++ {
		slice := data[i : i+4]

		var unique uint32
		for _, char := range slice {
			unique |= 1 << (int(char) - 97)
		}
		if bits.OnesCount32(unique) == 4 {
			return framework.TaskResult[int]{Value: i + 4, Error: nil}
		}
	}

	return
}

func Task2(filePath string) (result framework.TaskResult[int]) {
	_, err := framework.ReadFull(filePath)
	if err != nil {
		result.Error = err
		return
	}

	return
}
