// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day1 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day1.Task1()
            #expect(task != nil)
            let result = try task.example1()
            #expect("\(result)" == "11")
        }
        @Test
        func data() async throws {
            let task = Cases.Day1.Task1()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "1151792")
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day1.Task2()
            #expect(task != nil)
            let result = try task.example1()
            #expect("\(result)" == "31")
        }
        @Test
        func data() async throws {
            let task = Cases.Day1.Task2()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "21790168")
        }
    }
}
