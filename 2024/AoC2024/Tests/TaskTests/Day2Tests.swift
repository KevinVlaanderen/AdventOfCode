// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day2 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day2.Task1()
            try #expect(task.example1?.perform() == 2)
        }
        @Test
        func data() async throws {
            let task = Cases.Day2.Task1()
            try #expect(task.data?.perform() == 442)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day2.Task2()
            try #expect(task.example1?.perform() == 4)
        }
        @Test
        func data() async throws {
            let task = Cases.Day2.Task2()
            try #expect(task.data?.perform() == 493)
        }
    }
}
