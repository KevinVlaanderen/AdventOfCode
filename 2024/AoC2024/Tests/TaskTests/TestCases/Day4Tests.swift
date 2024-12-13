// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day4 {
    struct Task1 {
        @Test
        func example1() async throws {
            let day = Cases.Day4()
            let task = Cases.Day4.Task1()
            let param = task.param
            let expected = 18
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day4()
            let task = Cases.Day4.Task1()
            let param = task.param
            let expected = 2573
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let day = Cases.Day4()
            let task = Cases.Day4.Task2()
            let param = task.param
            let expected = 9
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day4()
            let task = Cases.Day4.Task2()
            let param = task.param
            let expected = 1850
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
}

import Testing
import Framework
@testable import Data

struct Day4 {
    struct Task1 {
        @Test
        func example1() async throws {
            let day = Cases.Day4()
            let task = Cases.Day4.Task1()
            let param = task.param
            let expected = 18
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day4()
            let task = Cases.Day4.Task1()
            let param = task.param
            let expected = 2573
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let day = Cases.Day4()
            let task = Cases.Day4.Task2()
            let param = task.param
            let expected = 9
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day4()
            let task = Cases.Day4.Task2()
            let param = task.param
            let expected = 1850
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
}
