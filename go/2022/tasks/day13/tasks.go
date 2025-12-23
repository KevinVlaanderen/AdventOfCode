//go:generate go run aoc/cmd/generate-tests

package day13

import (
	"aoc/2022/tasks/day13/model"
	"aoc/framework/tasks"
	"encoding/json"
	"go/types"
	"sort"

	"github.com/samber/lo"
)

// Task1 type:mock 	file:data	expected:13
// Task1 type:real 	file:day13 	expected:5623
func Task1(data string, _ types.Nil) (result tasks.Result[int]) {
	pairs := parse1(data)

	for _, pair := range pairs {
		if model.Compare(pair.Packet1, pair.Packet2) < 0 {
			result.Value += pair.Index
		}
	}

	return
}

// Task2 type:mock 	file:data	expected:140
// Task2 type:real 	file:day13 	expected:20570
func Task2(data string, _ types.Nil) (result tasks.Result[int]) {
	packets := parse2(data)

	divider1String := "[[2]]"
	divider2String := "[[6]]"
	var divider1, divider2 model.Packet
	_ = json.Unmarshal([]byte(divider1String), &divider1)
	_ = json.Unmarshal([]byte(divider2String), &divider2)

	packets = append(packets, divider1, divider2)

	sort.Slice(packets, func(i, j int) bool {
		return model.Compare(packets[i], packets[j]) < 0
	})

	var index1, index2 int
	for index, packet := range packets {
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

func parse1(data string) []model.Pair {
	pairIndex := 1
	blocks := tasks.LineBlocks(data)

	return lo.Map(blocks, func(block []string, index int) model.Pair {
		var packet1, packet2 model.Packet
		_ = json.Unmarshal([]byte(block[0]), &packet1)
		_ = json.Unmarshal([]byte(block[1]), &packet2)

		pair := model.Pair{Index: pairIndex, Packet1: packet1, Packet2: packet2}
		pairIndex++

		return pair
	})
}

func parse2(data string) []model.Packet {
	lines := lo.Flatten(tasks.LineBlocks(data))

	return lo.Map(lines, func(line string, index int) model.Packet {
		var packet model.Packet
		_ = json.Unmarshal([]byte(line), &packet)
		return packet
	})
}
