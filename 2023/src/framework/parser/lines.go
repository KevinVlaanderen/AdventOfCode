package parser

func Lines(line string) (value string, hasResult bool, err error) {
	if line != "" {
		value, hasResult = line, true
	}
	return value, hasResult, nil
}
