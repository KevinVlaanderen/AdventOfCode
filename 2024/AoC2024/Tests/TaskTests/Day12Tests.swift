// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day12 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day12.Task1()
            try #expect(task.example1?.perform() == 140)
        }
        @Test
        func example2() async throws {
            let task = Cases.Day12.Task1()
            try #expect(task.example2?.perform() == 772)
        }
        @Test
        func example3() async throws {
            let task = Cases.Day12.Task1()
            try #expect(task.example3?.perform() == 1930)
        }
        @Test
        func data() async throws {
            let task = Cases.Day12.Task1()
            try #expect(task.data?.perform() == 1457298)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day12.Task2()
            try #expect(task.example1?.perform() == 80)
        }
        @Test
        func example2() async throws {
            let task = Cases.Day12.Task2()
            try #expect(task.example2?.perform() == 436)
        }
        @Test
        func example4() async throws {
            let task = Cases.Day12.Task2()
            try #expect(task.example4?.perform() == 236)
        }
        @Test
        func example5() async throws {
            let task = Cases.Day12.Task2()
            try #expect(task.example5?.perform() == 368)
        }
        @Test
        func data() async throws {
            let task = Cases.Day12.Task2()
            try #expect(task.data?.perform() == 921636)
        }
    }
}
