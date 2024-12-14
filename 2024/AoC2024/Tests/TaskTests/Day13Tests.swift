// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day13 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day13.Task1()
            try #expect(task.example1?.perform() == 480)
        }
        @Test
        func data() async throws {
            let task = Cases.Day13.Task1()
            try #expect(task.data?.perform() == 35729)
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day13.Task2()
            try #expect(task.example1?.perform() == 875318608908)
        }
        @Test
        func data() async throws {
            let task = Cases.Day13.Task2()
            try #expect(task.data?.perform() == 88584689879723)
        }
    }
}
