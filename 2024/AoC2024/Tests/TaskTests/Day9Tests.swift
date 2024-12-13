// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day9 {
    struct Task1 {
        @Test
        func example1() async throws {
            let day = Cases.Day9()
            let task = Cases.Day9.Task1()
            let param = task.param
            let expected = 1928
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day9()
            let task = Cases.Day9.Task1()
            let param = task.param
            let expected = 6356833654075
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let day = Cases.Day9()
            let task = Cases.Day9.Task2()
            let param = task.param
            let expected = 2858
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day9()
            let task = Cases.Day9.Task2()
            let param = task.param
            let expected = 6389911791746
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
}
