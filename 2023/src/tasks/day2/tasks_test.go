package day2

import (
	"2023/src/framework"
	"2023/src/test"
	"github.com/samber/lo"
	"testing"
)

var taskDefinitions = []test.TaskDefinition[int]{
	{
		Task:     Task1,
		TestData: []test.TaskTestDefinition[int]{{"data", 8}},
		RealData: []test.TaskTestDefinition[int]{{"day2", 2278}},
	},
	{
		Task:     Task2,
		TestData: []test.TaskTestDefinition[int]{{"data", 2286}},
		RealData: []test.TaskTestDefinition[int]{{"day2", 67953}},
	},
}

func TestDay2(t *testing.T) {
	test.RunTests(t, taskDefinitions)
}

func BenchmarkDay2(b *testing.B) {
	test.RunBenchmarks(b, taskDefinitions)
}

func BenchmarkParse(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		total := 0
		for line := range framework.ReadLines(lo.Must(test.CreateRealDataPath("day2"))) {
			game := parse(line)
			total += game.id
		}
	}
}

func BenchmarkParseAll(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		games := parseAll(framework.ReadAll(lo.Must(test.CreateRealDataPath("day2"))))

		total := 0
		for _, game := range games {
			total += game.id
		}
	}
}

func BenchmarkParseAllChan(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		games := parseAllChan(framework.ReadAll(lo.Must(test.CreateRealDataPath("day2"))))

		total := 0
		for game := range games {
			total += game.id
		}
	}
}

func BenchmarkParseAllConcurrent(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		numWorkers := 4

		lineChannel := framework.ReadLines(lo.Must(test.CreateRealDataPath("day2")))
		inputChannels := lo.ChannelDispatcher(lineChannel, numWorkers, 0, lo.DispatchingStrategyRoundRobin[string])

		resultChannels := make([]<-chan int, numWorkers)

		for index := 0; index < numWorkers; index++ {
			resultChannels[index] = func(lines <-chan string) <-chan int {
				c := make(chan int)

				go func() {
					defer close(c)

					for line := range lines {
						game := parse(line)
						if game.valid() {
							c <- game.id
						}
					}
				}()

				return c
			}(inputChannels[index])
		}

		resultsChannel := lo.FanIn(0, resultChannels...)

		total := 0
		for result := range resultsChannel {
			total += result
		}
	}
}
