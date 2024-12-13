// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day5 {
    struct Task1 {
        @Test
        func example1() async throws {
            let day = Cases.Day5()
            let task = Cases.Day5.Task1()
            let param = task.param
            let expected = 143
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day5()
            let task = Cases.Day5.Task1()
            let param = task.param
            let expected = 4959
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let day = Cases.Day5()
            let task = Cases.Day5.Task2()
            let param = task.param
            let expected = 123
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day5()
            let task = Cases.Day5.Task2()
            let param = task.param
            let expected = 4655
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
}
