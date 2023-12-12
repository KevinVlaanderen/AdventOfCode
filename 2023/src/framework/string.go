package framework

func RepeatString(value string, count int, separator string) string {
	newValue := value
	for i := 1; i < count; i++ {
		newValue += separator + value
	}
	return newValue
}
