package framework

import (
	"github.com/samber/lo"
	"strings"
)

func CharLines(data string) [][]rune {
	return lo.Map(strings.Split(data, "\n"), func(item string, index int) []rune {
		return []rune(item)
	})
}

func Lines(data string) []string {
	return strings.Split(data, "\n")
}

func LineBlocks(data string) (blocks [][]string) {
	block := make([]string, 0)
	for _, line := range Lines(data) {
		if line == "" {
			if len(block) == 0 {
				continue
			}
			blocks = append(blocks, block)
			block = make([]string, 0)
		} else {
			block = append(block, line)
		}
	}
	if len(block) != 0 {
		blocks = append(blocks, block)
	}
	return blocks
}

func Blocks(data string) (blocks []string) {
	for _, block := range LineBlocks(data) {
		blocks = append(blocks, strings.Join(block, "\n"))
	}
	return blocks
}
