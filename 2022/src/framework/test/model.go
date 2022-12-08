package test

type Task[T comparable] func(fileName string) TaskResult[T]

type TaskResult[T comparable] struct {
	Value T
	Error error
}
