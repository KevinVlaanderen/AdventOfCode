// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day17 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day17.Task1()
            #expect(task != nil)
            let result = try task.example1()
            #expect("\(result)" == "4,6,3,5,6,3,5,2,1,0")
        }
        @Test
        func data() async throws {
            let task = Cases.Day17.Task1()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "7,3,1,3,6,3,6,0,2")
        }
    }
    struct Task2 {
        @Test
        func example2() async throws {
            let task = Cases.Day17.Task2()
            #expect(task != nil)
            let result = try task.example2()
            #expect("\(result)" == "117440")
        }
        @Test
        func data() async throws {
            let task = Cases.Day17.Task2()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "105843716614554")
        }
    }
}
