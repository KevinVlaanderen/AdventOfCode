// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day4 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day4.Task1()
            try #expect(task.example1?.perform() == 18)
        }
        @Test
        func data() async throws {
            let task = Cases.Day4.Task1()
            try #expect(task.data?.perform() == 2573)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day4.Task2()
            try #expect(task.example1?.perform() == 9)
        }
        @Test
        func data() async throws {
            let task = Cases.Day4.Task2()
            try #expect(task.data?.perform() == 1850)
        }
    }
}
