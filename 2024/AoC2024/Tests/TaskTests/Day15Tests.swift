// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day15 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day15.Task1()
            #expect(task != nil)
            let result = try task.example1()
            #expect("\(result)" == "10092")
        }
        @Test
        func example2() async throws {
            let task = Cases.Day15.Task1()
            #expect(task != nil)
            let result = try task.example2()
            #expect("\(result)" == "2028")
        }
        @Test
        func data() async throws {
            let task = Cases.Day15.Task1()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "1412971")
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day15.Task2()
            #expect(task != nil)
            let result = try task.example1()
            #expect("\(result)" == "9021")
        }
        @Test
        func data() async throws {
            let task = Cases.Day15.Task2()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "1429299")
        }
    }
}
