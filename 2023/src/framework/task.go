package framework

type Task[T comparable, P any] func(data string, param P) Result[T]

type Result[T comparable] struct {
	Value T
	Error error
}
