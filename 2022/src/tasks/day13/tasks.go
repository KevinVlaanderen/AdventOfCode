package day13

import (
	"2022/src/framework"
	"2022/src/framework/test"
	"encoding/json"
	"log"
	"sort"
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
	data, err := framework.ReadLineBlocks(filePath)
	if err != nil {
		result.Error = err
		return
	}

	divider1String := "[[2]]"
	divider2String := "[[6]]"
	var divider1, divider2 any
	_ = json.Unmarshal([]byte(divider1String), &divider1)
	_ = json.Unmarshal([]byte(divider2String), &divider2)

	allPackets := []any{divider1, divider2}

	for _, pair := range data {
		var packet1, packet2 any
		_ = json.Unmarshal([]byte(pair[0]), &packet1)
		_ = json.Unmarshal([]byte(pair[1]), &packet2)
		allPackets = append(allPackets, packet1, packet2)
	}

	sort.Slice(allPackets, func(i, j int) bool {
		return compare(allPackets[i], allPackets[j]) < 0
	})

	var index1, index2 int
	for index, packet := range allPackets {
		if packetBytes, err := json.Marshal(packet); err == nil {
			packetString := string(packetBytes)
			if packetString == divider1String {
				index1 = index + 1
			} else if packetString == divider2String {
				index2 = index + 1
			}
		}

	}
	result.Value = index1 * index2

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
