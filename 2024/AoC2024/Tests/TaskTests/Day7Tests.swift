// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day7 {
    struct Task1 {
        @Test
        func example1() async throws {
            let day = Cases.Day7()
            let task = Cases.Day7.Task1()
            let param = task.param
            let expected = 3749
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day7()
            let task = Cases.Day7.Task1()
            let param = task.param
            let expected = 5702958180383
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let day = Cases.Day7()
            let task = Cases.Day7.Task2()
            let param = task.param
            let expected = 11387
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day7()
            let task = Cases.Day7.Task2()
            let param = task.param
            let expected = 92612386119138
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
}
