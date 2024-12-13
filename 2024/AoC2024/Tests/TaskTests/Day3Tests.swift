// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day3 {
    struct Task1 {
        @Test
        func example1() async throws {
            let day = Cases.Day3()
            let task = Cases.Day3.Task1()
            let param = task.param
            let expected = 161
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day3()
            let task = Cases.Day3.Task1()
            let param = task.param
            let expected = 180233229
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let day = Cases.Day3()
            let task = Cases.Day3.Task2()
            let param = task.param
            let expected = 48
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day3()
            let task = Cases.Day3.Task2()
            let param = task.param
            let expected = 95411583
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
}
