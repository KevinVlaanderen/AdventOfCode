package framework

type Task[T comparable] func(data string) Result[T]

type Result[T comparable] struct {
	Value T
	Error error
}
