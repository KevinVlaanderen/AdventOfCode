// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day10 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day10.Task1()
            try #expect(task.example1?.perform() == 1)
        }
        @Test
        func example2() async throws {
            let task = Cases.Day10.Task1()
            try #expect(task.example2?.perform() == 2)
        }
        @Test
        func example3() async throws {
            let task = Cases.Day10.Task1()
            try #expect(task.example3?.perform() == 4)
        }
        @Test
        func example4() async throws {
            let task = Cases.Day10.Task1()
            try #expect(task.example4?.perform() == 3)
        }
        @Test
        func example5() async throws {
            let task = Cases.Day10.Task1()
            try #expect(task.example5?.perform() == 36)
        }
        @Test
        func data() async throws {
            let task = Cases.Day10.Task1()
            try #expect(task.data?.perform() == 816)
        }
    }
    struct Task2 {
        @Test
        func example6() async throws {
            let task = Cases.Day10.Task2()
            try #expect(task.example6?.perform() == 3)
        }
        @Test
        func example7() async throws {
            let task = Cases.Day10.Task2()
            try #expect(task.example7?.perform() == 13)
        }
        @Test
        func example8() async throws {
            let task = Cases.Day10.Task2()
            try #expect(task.example8?.perform() == 227)
        }
        @Test
        func example5() async throws {
            let task = Cases.Day10.Task2()
            try #expect(task.example5?.perform() == 81)
        }
        @Test
        func data() async throws {
            let task = Cases.Day10.Task2()
            try #expect(task.data?.perform() == 1960)
        }
    }
}
