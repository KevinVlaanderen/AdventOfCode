// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day16 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day16.Task1()
            try #expect(task.example1?.perform() == 7036)
        }
        @Test
        func example2() async throws {
            let task = Cases.Day16.Task1()
            try #expect(task.example2?.perform() == 11048)
        }
        @Test
        func data() async throws {
            let task = Cases.Day16.Task1()
            try #expect(task.data?.perform() == 143580)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day16.Task2()
            try #expect(task.example1?.perform() == 45)
        }
        @Test
        func example2() async throws {
            let task = Cases.Day16.Task2()
            try #expect(task.example2?.perform() == 64)
        }
        @Test
        func data() async throws {
            let task = Cases.Day16.Task2()
            try #expect(task.data?.perform() == 645)
        }
    }
}
