// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day2 {
    struct Task1 {
        @Test
        func example1() async throws {
            let day = Cases.Day2()
            let task = Cases.Day2.Task1()
            let param = task.param
            let expected = 2
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day2()
            let task = Cases.Day2.Task1()
            let param = task.param
            let expected = 442
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let day = Cases.Day2()
            let task = Cases.Day2.Task2()
            let param = task.param
            let expected = 4
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day2()
            let task = Cases.Day2.Task2()
            let param = task.param
            let expected = 493
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
}

import Testing
import Framework
@testable import Data

struct Day2 {
    struct Task1 {
        @Test
        func example1() async throws {
            let day = Cases.Day2()
            let task = Cases.Day2.Task1()
            let param = task.param
            let expected = 2
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day2()
            let task = Cases.Day2.Task1()
            let param = task.param
            let expected = 442
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let day = Cases.Day2()
            let task = Cases.Day2.Task2()
            let param = task.param
            let expected = 4
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day2()
            let task = Cases.Day2.Task2()
            let param = task.param
            let expected = 493
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
}
