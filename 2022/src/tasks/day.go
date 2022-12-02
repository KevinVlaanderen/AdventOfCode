package tasks

type Day interface {
	Task1(fileName string) (*int, error)
	Task2(fileName string) (*int, error)
}
