// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day11 {
    struct Task1 {
        @Test
        func example1() async throws {
            let day = Cases.Day11()
            let task = Cases.Day11.Task1()
            let param = task.param
            let expected = 55312
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day11()
            let task = Cases.Day11.Task1()
            let param = task.param
            let expected = 213625
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
    struct Task2 {
        @Test
        func data() async throws {
            let day = Cases.Day11()
            let task = Cases.Day11.Task2()
            let param = task.param
            let expected = 252442982856820
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
}
