// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day12 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day12.Task1()
            #expect(task != nil)
            let result = try task.example1!.perform()
            #expect("\(result)" == "140")
        }
        @Test
        func example2() async throws {
            let task = Cases.Day12.Task1()
            #expect(task != nil)
            let result = try task.example2!.perform()
            #expect("\(result)" == "772")
        }
        @Test
        func example3() async throws {
            let task = Cases.Day12.Task1()
            #expect(task != nil)
            let result = try task.example3!.perform()
            #expect("\(result)" == "1930")
        }
        @Test
        func data() async throws {
            let task = Cases.Day12.Task1()
            #expect(task != nil)
            let result = try task.data!.perform()
            #expect("\(result)" == "1457298")
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day12.Task2()
            #expect(task != nil)
            let result = try task.example1!.perform()
            #expect("\(result)" == "80")
        }
        @Test
        func example2() async throws {
            let task = Cases.Day12.Task2()
            #expect(task != nil)
            let result = try task.example2!.perform()
            #expect("\(result)" == "436")
        }
        @Test
        func example4() async throws {
            let task = Cases.Day12.Task2()
            #expect(task != nil)
            let result = try task.example4!.perform()
            #expect("\(result)" == "236")
        }
        @Test
        func example5() async throws {
            let task = Cases.Day12.Task2()
            #expect(task != nil)
            let result = try task.example5!.perform()
            #expect("\(result)" == "368")
        }
        @Test
        func data() async throws {
            let task = Cases.Day12.Task2()
            #expect(task != nil)
            let result = try task.data!.perform()
            #expect("\(result)" == "921636")
        }
    }
}
