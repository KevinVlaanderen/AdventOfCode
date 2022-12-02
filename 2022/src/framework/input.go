package framework

import (
	"bufio"
	"os"
)

func ReadLineBlocks(filePath string) (result [][]string, err error) {
	f, err := os.Open(filePath)

	defer func(f *os.File) {
		if tempErr := f.Close(); err == nil {
			err = tempErr
		}
	}(f)

	scanner := bufio.NewScanner(f)

	var block []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			result = append(result, block)
			block = nil
		} else {
			block = append(block, line)
		}
	}
	result = append(result, block)

	return
}
