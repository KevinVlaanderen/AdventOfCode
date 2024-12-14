// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day3 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day3.Task1()
            try #expect(task.example1?.perform() == 161)
        }
        @Test
        func data() async throws {
            let task = Cases.Day3.Task1()
            try #expect(task.data?.perform() == 180233229)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day3.Task2()
            try #expect(task.example1?.perform() == 48)
        }
        @Test
        func data() async throws {
            let task = Cases.Day3.Task2()
            try #expect(task.data?.perform() == 95411583)
        }
    }
}
