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
            #expect(task != nil)
            let result = try task.example1()
            #expect("\(result)" == "7036")
        }
        @Test
        func example2() async throws {
            let task = Cases.Day16.Task1()
            #expect(task != nil)
            let result = try task.example2()
            #expect("\(result)" == "11048")
        }
        @Test
        func data() async throws {
            let task = Cases.Day16.Task1()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "143580")
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day16.Task2()
            #expect(task != nil)
            let result = try task.example1()
            #expect("\(result)" == "45")
        }
        @Test
        func example2() async throws {
            let task = Cases.Day16.Task2()
            #expect(task != nil)
            let result = try task.example2()
            #expect("\(result)" == "64")
        }
        @Test
        func data() async throws {
            let task = Cases.Day16.Task2()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "645")
        }
    }
}
