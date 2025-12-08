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
            #expect(task != nil)
            let result = try task.example1()
            #expect("\(result)" == "480")
        }
        @Test
        func data() async throws {
            let task = Cases.Day13.Task1()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "35729")
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day13.Task2()
            #expect(task != nil)
            let result = try task.example1()
            #expect("\(result)" == "875318608908")
        }
        @Test
        func data() async throws {
            let task = Cases.Day13.Task2()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "88584689879723")
        }
    }
}
