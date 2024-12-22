// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day11 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day11.Task1()
            #expect(task != nil)
            let result = try task.example1()
            #expect("\(result)" == "55312")
        }
        @Test
        func data() async throws {
            let task = Cases.Day11.Task1()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "213625")
        }
    }
    struct Task2 {
        @Test
        func data() async throws {
            let task = Cases.Day11.Task2()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "252442982856820")
        }
    }
}
