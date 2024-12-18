// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day18 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day18.Task1()
            #expect(task != nil)
            let result = try task.example1!.perform()
            #expect("\(result)" == "22")
        }
        @Test
        func data() async throws {
            let task = Cases.Day18.Task1()
            #expect(task != nil)
            let result = try task.data!.perform()
            #expect("\(result)" == "270")
        }
    }
}
