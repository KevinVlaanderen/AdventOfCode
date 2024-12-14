// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day8 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day8.Task1()
            try #expect(task.example1?.perform() == 14)
        }
        @Test
        func data() async throws {
            let task = Cases.Day8.Task1()
            try #expect(task.data?.perform() == 320)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day8.Task2()
            try #expect(task.example1?.perform() == 34)
        }
        @Test
        func data() async throws {
            let task = Cases.Day8.Task2()
            try #expect(task.data?.perform() == 1157)
        }
    }
}
