// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day7 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day7.Task1()
            try #expect(task.example1?.perform() == 3749)
        }
        @Test
        func data() async throws {
            let task = Cases.Day7.Task1()
            try #expect(task.data?.perform() == 5702958180383)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day7.Task2()
            try #expect(task.example1?.perform() == 11387)
        }
        @Test
        func data() async throws {
            let task = Cases.Day7.Task2()
            try #expect(task.data?.perform() == 92612386119138)
        }
    }
}
