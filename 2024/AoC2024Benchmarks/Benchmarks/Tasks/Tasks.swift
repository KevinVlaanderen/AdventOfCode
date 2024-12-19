// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT

import Benchmark
import Framework
@testable import Data

@MainActor
let benchmarks = {
    Benchmark("Day1.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day1.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day1.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day1.Task2()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day2.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day2.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day2.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day2.Task2()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day3.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day3.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day3.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day3.Task2()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day4.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day4.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day4.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day4.Task2()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day5.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day5.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day5.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day5.Task2()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day6.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day6.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day6.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day6.Task2()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day7.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day7.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day7.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day7.Task2()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day8.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day8.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day8.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day8.Task2()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day9.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day9.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day9.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day9.Task2()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day10.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day10.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day10.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day10.Task2()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day11.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day11.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day11.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day11.Task2()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day12.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day12.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day12.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day12.Task2()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day13.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day13.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day13.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day13.Task2()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day14.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day14.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day14.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day14.Task2()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day15.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day15.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day15.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day15.Task2()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day16.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day16.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day16.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day16.Task2()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day17.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day17.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day18.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day18.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day18.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day18.Task2()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day19.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day19.Task1()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
    Benchmark("Day19.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let task = Cases.Day19.Task2()
        for _ in benchmark.scaledIterations {
            try blackHole(task.data!.perform())
        }
    }
}
