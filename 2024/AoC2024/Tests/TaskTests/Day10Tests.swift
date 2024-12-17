// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day10 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day10.Task1()
            #expect(task != nil)
            let result = try task.example1!.perform()
            #expect("\(result)" == "1")
        }
        @Test
        func example2() async throws {
            let task = Cases.Day10.Task1()
            #expect(task != nil)
            let result = try task.example2!.perform()
            #expect("\(result)" == "2")
        }
        @Test
        func example3() async throws {
            let task = Cases.Day10.Task1()
            #expect(task != nil)
            let result = try task.example3!.perform()
            #expect("\(result)" == "4")
        }
        @Test
        func example4() async throws {
            let task = Cases.Day10.Task1()
            #expect(task != nil)
            let result = try task.example4!.perform()
            #expect("\(result)" == "3")
        }
        @Test
        func example5() async throws {
            let task = Cases.Day10.Task1()
            #expect(task != nil)
            let result = try task.example5!.perform()
            #expect("\(result)" == "36")
        }
        @Test
        func data() async throws {
            let task = Cases.Day10.Task1()
            #expect(task != nil)
            let result = try task.data!.perform()
            #expect("\(result)" == "816")
        }
    }
    struct Task2 {
        @Test
        func example6() async throws {
            let task = Cases.Day10.Task2()
            #expect(task != nil)
            let result = try task.example6!.perform()
            #expect("\(result)" == "3")
        }
        @Test
        func example7() async throws {
            let task = Cases.Day10.Task2()
            #expect(task != nil)
            let result = try task.example7!.perform()
            #expect("\(result)" == "13")
        }
        @Test
        func example8() async throws {
            let task = Cases.Day10.Task2()
            #expect(task != nil)
            let result = try task.example8!.perform()
            #expect("\(result)" == "227")
        }
        @Test
        func example5() async throws {
            let task = Cases.Day10.Task2()
            #expect(task != nil)
            let result = try task.example5!.perform()
            #expect("\(result)" == "81")
        }
        @Test
        func data() async throws {
            let task = Cases.Day10.Task2()
            #expect(task != nil)
            let result = try task.data!.perform()
            #expect("\(result)" == "1960")
        }
    }
}
