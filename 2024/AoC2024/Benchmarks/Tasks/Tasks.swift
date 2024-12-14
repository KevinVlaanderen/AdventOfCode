// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT

import Benchmark
import Framework
@testable import Data

@MainActor
let benchmarks = {
    Benchmark("Cases.Day1.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day1()
        let task = Cases.Day1.Task1()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day1.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day1()
        let task = Cases.Day1.Task2()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day2.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day2()
        let task = Cases.Day2.Task1()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day2.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day2()
        let task = Cases.Day2.Task2()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day3.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day3()
        let task = Cases.Day3.Task1()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day3.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day3()
        let task = Cases.Day3.Task2()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day4.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day4()
        let task = Cases.Day4.Task1()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day4.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day4()
        let task = Cases.Day4.Task2()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day5.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day5()
        let task = Cases.Day5.Task1()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day5.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day5()
        let task = Cases.Day5.Task2()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day6.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day6()
        let task = Cases.Day6.Task1()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day6.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day6()
        let task = Cases.Day6.Task2()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day7.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day7()
        let task = Cases.Day7.Task1()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day7.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day7()
        let task = Cases.Day7.Task2()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day8.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day8()
        let task = Cases.Day8.Task1()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day8.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day8()
        let task = Cases.Day8.Task2()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day9.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day9()
        let task = Cases.Day9.Task1()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day9.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day9()
        let task = Cases.Day9.Task2()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day10.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day10()
        let task = Cases.Day10.Task1()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day10.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day10()
        let task = Cases.Day10.Task2()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day11.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day11()
        let task = Cases.Day11.Task1()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day11.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day11()
        let task = Cases.Day11.Task2()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day12.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day12()
        let task = Cases.Day12.Task1()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day12.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day12()
        let task = Cases.Day12.Task2()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day13.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day13()
        let task = Cases.Day13.Task1()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day13.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day13()
        let task = Cases.Day13.Task2()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day14.Task1.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day14()
        let task = Cases.Day14.Task1()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
    Benchmark("Cases.Day14.Task2.data") { benchmark in
        benchmark.configuration.timeUnits = .microseconds
        let day = Cases.Day14()
        let task = Cases.Day14.Task2()
        let param = task.param
        for _ in benchmark.scaledIterations {
            blackHole(try execute(day: day.day, task: task.task, data: task.data, param: param))
        }
    }
}
