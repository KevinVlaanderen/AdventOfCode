// Generated using Sourcery 2.2.5 — https://github.com/krzysztofzablocki/Sourcery
// DO NOT EDIT
import Testing
import Framework
@testable import Data

struct Day23 {
    struct Task1 {
        @Test
        func example1() async throws {
            let task = Cases.Day23.Task1()
            #expect(task != nil)
            let result = try task.example1()
            #expect("\(result)" == "7")
        }
        @Test
        func data() async throws {
            let task = Cases.Day23.Task1()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "1352")
        }
    }
    struct Task2 {
        @Test
        func example1() async throws {
            let task = Cases.Day23.Task2()
            #expect(task != nil)
            let result = try task.example1()
            #expect("\(result)" == "co,de,ka,ta")
        }
        @Test
        func data() async throws {
            let task = Cases.Day23.Task2()
            #expect(task != nil)
            let result = try task.data()
            #expect("\(result)" == "dm,do,fr,gf,gh,gy,iq,jb,kt,on,rg,xf,ze")
        }
    }
}
