package day13

import (
	"2022/src/framework/tasks"
	"2022/src/tasks/day13/model"
	"encoding/json"
	"sort"
)

func Task1(filePath string) (result tasks.TaskResult[int]) {
	pairs := tasks.ReadStream(filePath, createTask1Parser())

	for pair := range pairs {
		if model.Compare(pair.Packet1, pair.Packet2) < 0 {
			result.Value += pair.Index
		}
	}

	return
}

func Task2(filePath string) (result tasks.TaskResult[int]) {
	packets := tasks.Read(filePath, task2Parser)

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

func createTask1Parser() func(line string) (result model.Pair, hasResult bool, err error) {
	index := 1
	var lines []string

	return func(line string) (result model.Pair, hasResult bool, err error) {
		if line != "" {
			lines = append(lines, line)
			return result, false, err
		} else if line == "" && len(lines) == 0 {
			return result, false, err
		}

		var packet1, packet2 model.Packet
		_ = json.Unmarshal([]byte(lines[0]), &packet1)
		_ = json.Unmarshal([]byte(lines[1]), &packet2)

		pair := model.Pair{Index: index, Packet1: packet1, Packet2: packet2}

		index++
		lines = nil

		return pair, true, nil
	}
}

func task2Parser(line string) (result model.Packet, hasResult bool, err error) {
	if line == "" {
		return result, false, err
	}

	var packet model.Packet
	_ = json.Unmarshal([]byte(line), &packet)

	return packet, true, nil
}