package day7

type InputType int

const (
	COMMAND InputType = iota
	CONTENT
)

type Input struct {
	InputType
	data string
}

func Parse(data []string) (inputs []Input) {
	for _, line := range data {
		if line == "" {
			continue
		}

		if line[0] == '$' {
			inputs = append(inputs, Input{
				InputType: COMMAND,
				data:      line[2:],
			})
		} else {
			inputs = append(inputs, Input{
				InputType: CONTENT,
				data:      line,
			})
		}
	}
	return
}
