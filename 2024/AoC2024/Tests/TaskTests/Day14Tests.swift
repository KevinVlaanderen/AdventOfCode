// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day14 {
    struct Task1 {
        @Test
        func example1() async throws {
            let day = Cases.Day14()
            let task = Cases.Day14.Task1()
            let param = task.param
            let expected = 12
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day14()
            let task = Cases.Day14.Task1()
            let param = task.param
            let expected = 208437768
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
}
