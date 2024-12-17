// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day7 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day7.Task1()
            #expect(task != nil)
            let result = try task.example1!.perform()
            #expect("\(result)" == "3749")
        }
        @Test
        func data() async throws {
            let task = Cases.Day7.Task1()
            #expect(task != nil)
            let result = try task.data!.perform()
            #expect("\(result)" == "5702958180383")
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day7.Task2()
            #expect(task != nil)
            let result = try task.example1!.perform()
            #expect("\(result)" == "11387")
        }
        @Test
        func data() async throws {
            let task = Cases.Day7.Task2()
            #expect(task != nil)
            let result = try task.data!.perform()
            #expect("\(result)" == "92612386119138")
        }
    }
}
