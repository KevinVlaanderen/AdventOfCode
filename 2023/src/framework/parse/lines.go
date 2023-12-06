package parse

import "2023/src/framework/task"

func Lines() task.Parser[string] {
	return func(line string) (value string, hasResult bool, err error) {
		if line != "" {
			value, hasResult = line, true
		}
		return value, hasResult, nil
	}
}
