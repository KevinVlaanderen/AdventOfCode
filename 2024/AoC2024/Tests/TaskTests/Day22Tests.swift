// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day22 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day22.Task1()
            #expect(task != nil)
            let result = try task.example1()
            #expect("\(result)" == "37327623")
        }
        @Test
        func data() async throws {
            let task = Cases.Day22.Task1()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "15006633487")
        }
    }
    struct Task2 {
        @Test
        func example2() async throws {
            let task = Cases.Day22.Task2()
            #expect(task != nil)
            let result = try task.example2()
            #expect("\(result)" == "23")
        }
        @Test
        func data() async throws {
            let task = Cases.Day22.Task2()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "1710")
        }
    }
}
