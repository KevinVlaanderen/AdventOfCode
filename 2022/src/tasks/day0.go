package tasks

type Day0 struct {
}

func (d Day0) task1(data string) int {
	if data == "day0" {
		return 1
	} else {
		return 0
	}
}
