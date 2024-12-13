// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day1 {
    struct Task1 {
        @Test
        func example1() async throws {
            let day = Cases.Day1()
            let task = Cases.Day1.Task1()
            let param = task.param
            let expected = 11
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day1()
            let task = Cases.Day1.Task1()
            let param = task.param
            let expected = 1151792
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let day = Cases.Day1()
            let task = Cases.Day1.Task2()
            let param = task.param
            let expected = 31
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day1()
            let task = Cases.Day1.Task2()
            let param = task.param
            let expected = 21790168
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
}
