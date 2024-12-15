// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day15 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day15.Task1()
            try #expect(task.example1?.perform() == 10092)
        }
        @Test
        func example2() async throws {
            let task = Cases.Day15.Task1()
            try #expect(task.example2?.perform() == 2028)
        }
        @Test
        func data() async throws {
            let task = Cases.Day15.Task1()
            try #expect(task.data?.perform() == 1412971)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day15.Task2()
            try #expect(task.example1?.perform() == 9021)
        }
        @Test
        func data() async throws {
            let task = Cases.Day15.Task2()
            try #expect(task.data?.perform() == 1429299)
        }
    }
}
