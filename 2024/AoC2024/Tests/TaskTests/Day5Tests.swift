// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day5 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day5.Task1()
            try #expect(task.example1?.perform() == 143)
        }
        @Test
        func data() async throws {
            let task = Cases.Day5.Task1()
            try #expect(task.data?.perform() == 4959)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day5.Task2()
            try #expect(task.example1?.perform() == 123)
        }
        @Test
        func data() async throws {
            let task = Cases.Day5.Task2()
            try #expect(task.data?.perform() == 4655)
        }
    }
}
