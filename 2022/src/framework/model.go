package framework

type Task = func(fileName string) TaskResult

type TaskResult struct {
	Value int
	Error error
}
