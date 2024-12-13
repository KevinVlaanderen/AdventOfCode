// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day8 {
    struct Task1 {
        @Test
        func example1() async throws {
            let day = Cases.Day8()
            let task = Cases.Day8.Task1()
            let param = task.param
            let expected = 14
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day8()
            let task = Cases.Day8.Task1()
            let param = task.param
            let expected = 320
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let day = Cases.Day8()
            let task = Cases.Day8.Task2()
            let param = task.param
            let expected = 34
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day8()
            let task = Cases.Day8.Task2()
            let param = task.param
            let expected = 1157
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
}
