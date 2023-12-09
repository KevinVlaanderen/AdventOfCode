package framework

import (
	"bufio"
	"log"
	"os"
)

type Parser[T any] func(line string) (result T, hasResult bool, err error)

func ReadStream[T any](
	filePath string,
	parser Parser[T],
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
	parser Parser[T],
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

func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			log.Panic(err)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
