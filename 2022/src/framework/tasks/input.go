package tasks

import (
	"bufio"
	"log"
	"os"
)

func ReadStream[T any](
	filePath string,
	parser func(line string) (result T, hasResult bool, err error),
) (data <-chan T) {
	outputChannel := make(chan T)

	go func() {
		defer close(outputChannel)

		f, err := os.Open(filePath)
		if err != nil {
			log.Panic(err)
		}

		defer func(f *os.File) {
			if err = f.Close(); err != nil {
				log.Panic(err)
			}
		}(f)

		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)

		handler := func(result T, hasResult bool, err error) {
			switch {
			case err != nil:
				log.Panic(err)
			case hasResult:
				outputChannel <- result
			}
		}

		for scanner.Scan() {
			if scanner.Err() != nil {
				log.Panic(scanner.Err())
			}
			handler(parser(scanner.Text()))
		}
		handler(parser(""))
	}()

	return outputChannel
}

func Read[T any](
	filePath string,
	parser func(line string) (result T, hasResult bool, err error),
) (output []T) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Panic(err)
	}

	defer func(f *os.File) {
		if err = f.Close(); err != nil {
			log.Panic(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	handler := func(result T, hasResult bool, err error) {
		switch {
		case err != nil:
			log.Panic(err)
		case hasResult:
			output = append(output, result)
		}
	}

	for scanner.Scan() {
		if scanner.Err() != nil {
			log.Panic(scanner.Err())
		}
		handler(parser(scanner.Text()))
	}
	handler(parser(""))

	return
}
