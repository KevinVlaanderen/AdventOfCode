package parse

import "2023/src/framework/task"

func Blocks() task.Parser[[]string] {
	lines := make([]string, 0)
	return func(line string) (value []string, hasResult bool, err error) {
		if line == "" {
			result := lines
			lines = make([]string, 0)
			return result, true, nil
		}
		lines = append(lines, line)
		return value, hasResult, nil
	}
}
