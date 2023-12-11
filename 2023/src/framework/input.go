package framework

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Parser[T any] func(line string) (result T, hasResult bool, err error)

func ParseLines[T any](
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

func ParseAllLines[T any](
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

func ReadAll(path string) string {
	if bytes, err := os.ReadFile(path); err != nil {
		panic(err)
	} else {
		return string(bytes)
	}
}

func ReadAllLines(path string) (lines []string) {
	file := MustOpenFile(path)
	defer MustCloseFile(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return
}

func ReadLines(path string) <-chan string {
	c := make(chan string)

	go func() {
		defer close(c)

		file := MustOpenFile(path)
		defer MustCloseFile(file)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			c <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}()

	return c
}

func ReadBlocks(path string) <-chan string {
	c := make(chan string)

	go func() {
		defer close(c)

		file := MustOpenFile(path)
		defer MustCloseFile(file)

		lines := make([]string, 0)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if line := scanner.Text(); line == "" && len(lines) > 0 {
				c <- strings.Join(lines, "\n")
				lines = make([]string, 0)
			} else {
				lines = append(lines, line)
			}
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		if len(lines) > 0 {
			c <- strings.Join(lines, "\n")
		}
	}()

	return c
}

func MustOpenFile(path string) *os.File {
	if file, err := os.Open(path); err != nil {
		log.Panic(err)
		return file
	} else {
		return file
	}
}

func MustCloseFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Panic(err)
	}
}
