// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day6 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day6.Task1()
            #expect(task != nil)
            let result = try task.example1()
            #expect("\(result)" == "41")
        }
        @Test
        func data() async throws {
            let task = Cases.Day6.Task1()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "4973")
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day6.Task2()
            #expect(task != nil)
            let result = try task.example1()
            #expect("\(result)" == "6")
        }
        @Test
        func data() async throws {
            let task = Cases.Day6.Task2()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "1482")
        }
    }
}
