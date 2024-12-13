// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day10 {
    struct Task1 {
        @Test
        func example1() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task1()
            let param = task.param
            let expected = 1
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func example2() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task1()
            let param = task.param
            let expected = 2
            try #expect(execute(day: day.day, task: task.task, data: task.example2, param: param) == expected)
        }
        @Test
        func example3() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task1()
            let param = task.param
            let expected = 4
            try #expect(execute(day: day.day, task: task.task, data: task.example3, param: param) == expected)
        }
        @Test
        func example4() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task1()
            let param = task.param
            let expected = 3
            try #expect(execute(day: day.day, task: task.task, data: task.example4, param: param) == expected)
        }
        @Test
        func example5() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task1()
            let param = task.param
            let expected = 36
            try #expect(execute(day: day.day, task: task.task, data: task.example5, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task1()
            let param = task.param
            let expected = 816
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
    struct Task2 {
        @Test
        func example6() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task2()
            let param = task.param
            let expected = 3
            try #expect(execute(day: day.day, task: task.task, data: task.example6, param: param) == expected)
        }
        @Test
        func example7() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task2()
            let param = task.param
            let expected = 13
            try #expect(execute(day: day.day, task: task.task, data: task.example7, param: param) == expected)
        }
        @Test
        func example8() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task2()
            let param = task.param
            let expected = 227
            try #expect(execute(day: day.day, task: task.task, data: task.example8, param: param) == expected)
        }
        @Test
        func example5() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task2()
            let param = task.param
            let expected = 81
            try #expect(execute(day: day.day, task: task.task, data: task.example5, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task2()
            let param = task.param
            let expected = 1960
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
}

import Testing
import Framework
@testable import Data

struct Day10 {
    struct Task1 {
        @Test
        func example1() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task1()
            let param = task.param
            let expected = 1
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func example2() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task1()
            let param = task.param
            let expected = 2
            try #expect(execute(day: day.day, task: task.task, data: task.example2, param: param) == expected)
        }
        @Test
        func example3() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task1()
            let param = task.param
            let expected = 4
            try #expect(execute(day: day.day, task: task.task, data: task.example3, param: param) == expected)
        }
        @Test
        func example4() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task1()
            let param = task.param
            let expected = 3
            try #expect(execute(day: day.day, task: task.task, data: task.example4, param: param) == expected)
        }
        @Test
        func example5() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task1()
            let param = task.param
            let expected = 36
            try #expect(execute(day: day.day, task: task.task, data: task.example5, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task1()
            let param = task.param
            let expected = 816
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
    struct Task2 {
        @Test
        func example6() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task2()
            let param = task.param
            let expected = 3
            try #expect(execute(day: day.day, task: task.task, data: task.example6, param: param) == expected)
        }
        @Test
        func example7() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task2()
            let param = task.param
            let expected = 13
            try #expect(execute(day: day.day, task: task.task, data: task.example7, param: param) == expected)
        }
        @Test
        func example8() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task2()
            let param = task.param
            let expected = 227
            try #expect(execute(day: day.day, task: task.task, data: task.example8, param: param) == expected)
        }
        @Test
        func example5() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task2()
            let param = task.param
            let expected = 81
            try #expect(execute(day: day.day, task: task.task, data: task.example5, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day10()
            let task = Cases.Day10.Task2()
            let param = task.param
            let expected = 1960
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
}
