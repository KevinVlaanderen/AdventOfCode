// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day12 {
    struct Task1 {
        @Test
        func example1() async throws {
            let day = Cases.Day12()
            let task = Cases.Day12.Task1()
            let param = task.param
            let expected = 140
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func example2() async throws {
            let day = Cases.Day12()
            let task = Cases.Day12.Task1()
            let param = task.param
            let expected = 772
            try #expect(execute(day: day.day, task: task.task, data: task.example2, param: param) == expected)
        }
        @Test
        func example3() async throws {
            let day = Cases.Day12()
            let task = Cases.Day12.Task1()
            let param = task.param
            let expected = 1930
            try #expect(execute(day: day.day, task: task.task, data: task.example3, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day12()
            let task = Cases.Day12.Task1()
            let param = task.param
            let expected = 1457298
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let day = Cases.Day12()
            let task = Cases.Day12.Task2()
            let param = task.param
            let expected = 80
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func example2() async throws {
            let day = Cases.Day12()
            let task = Cases.Day12.Task2()
            let param = task.param
            let expected = 436
            try #expect(execute(day: day.day, task: task.task, data: task.example2, param: param) == expected)
        }
        @Test
        func example4() async throws {
            let day = Cases.Day12()
            let task = Cases.Day12.Task2()
            let param = task.param
            let expected = 236
            try #expect(execute(day: day.day, task: task.task, data: task.example4, param: param) == expected)
        }
        @Test
        func example5() async throws {
            let day = Cases.Day12()
            let task = Cases.Day12.Task2()
            let param = task.param
            let expected = 368
            try #expect(execute(day: day.day, task: task.task, data: task.example5, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day12()
            let task = Cases.Day12.Task2()
            let param = task.param
            let expected = 921636
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
}
