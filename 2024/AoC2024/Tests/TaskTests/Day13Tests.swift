// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day13 {
    struct Task1 {
        @Test
        func example1() async throws {
            let day = Cases.Day13()
            let task = Cases.Day13.Task1()
            let param = task.param
            let expected = 480
            try #expect(execute(day: day.day, task: task.task, data: task.example1, param: param) == expected)
        }
        @Test
        func data() async throws {
            let day = Cases.Day13()
            let task = Cases.Day13.Task1()
            let param = task.param
            let expected = -1
            try #expect(execute(day: day.day, task: task.task, data: task.data, param: param) == expected)
        }
    }
}
