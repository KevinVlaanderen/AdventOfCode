package day13

import (
	"2022/src/framework"
	"2022/src/framework/test"
	"encoding/json"
	"log"
)

func Task1(filePath string) (result test.TaskResult[int]) {
	data, err := framework.ReadLineBlocks(filePath)
	if err != nil {
		result.Error = err
		return
	}

	for index, pair := range data {
		var packet1, packet2 any
		_ = json.Unmarshal([]byte(pair[0]), &packet1)
		_ = json.Unmarshal([]byte(pair[1]), &packet2)

		if compare(packet1, packet2) < 0 {
			log.Printf("pair %v is in the right order\n", index+1)
			result.Value += index + 1
		}
	}

	return
}

func Task2(filePath string) (result test.TaskResult[int]) {
	_, err := framework.ReadLineBlocks(filePath)
	if err != nil {
		result.Error = err
		return
	}

	return
}

func compare(a any, b any) int {
	aSlice, aIsList := a.([]any)
	bSlice, bIsList := b.([]any)

	if !aIsList && !bIsList {
		return int(a.(float64) - b.(float64))
	} else {
		if aIsList && !bIsList {
			bSlice = []any{b}
		} else if !aIsList && bIsList {
			aSlice = []any{a}
		}

		if len(aSlice) == 0 && len(bSlice) == 0 {
			return 0
		} else if len(aSlice) == 0 {
			return -1
		} else if len(bSlice) == 0 {
			return 1
		}

		for index := range aSlice {
			if index >= len(bSlice) {
				return 1
			}

			itemA := aSlice[index]
			itemB := bSlice[index]

			switch result := compare(itemA, itemB); {
			case result > 0:
				return 1
			case result < 0:
				return -1
			}
		}

		if len(bSlice) > len(aSlice) {
			return -1
		}
	}

	return 0
}
