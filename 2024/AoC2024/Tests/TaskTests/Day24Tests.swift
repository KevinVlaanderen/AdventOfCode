// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day24 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day24.Task1()
            #expect(task != nil)
            let result = try task.example1()
            #expect("\(result)" == "4")
        }
        @Test
        func example2() async throws {
            let task = Cases.Day24.Task1()
            #expect(task != nil)
            let result = try task.example2()
            #expect("\(result)" == "2024")
        }
        @Test
        func data() async throws {
            let task = Cases.Day24.Task1()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "60714423975686")
        }
    }
    struct Task2 {
    }
}
