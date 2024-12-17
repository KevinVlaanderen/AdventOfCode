// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day9 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day9.Task1()
            #expect(task != nil)
            let result = try task.example1!.perform()
            #expect("\(result)" == "1928")
        }
        @Test
        func data() async throws {
            let task = Cases.Day9.Task1()
            #expect(task != nil)
            let result = try task.data!.perform()
            #expect("\(result)" == "6356833654075")
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day9.Task2()
            #expect(task != nil)
            let result = try task.example1!.perform()
            #expect("\(result)" == "2858")
        }
        @Test
        func data() async throws {
            let task = Cases.Day9.Task2()
            #expect(task != nil)
            let result = try task.data!.perform()
            #expect("\(result)" == "6389911791746")
        }
    }
}
