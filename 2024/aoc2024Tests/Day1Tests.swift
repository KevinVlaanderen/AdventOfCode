import Testing
@testable import aoc2024

final class Day1Tests {
    @Test(arguments: [
        (Day1.example1, (), 11),
        (Day1.data, (), 1151792)
    ])
    func task1(data: String, param: Day1.P1, result: Day1.R1) throws {
        try #expect(Day1.task1(data: data, param: param) == result)
    }
    
    @Test(arguments: [
        (Day1.example1, (), 31),
        (Day1.data, (), 21790168)
    ])
    func task2(data: String, param: Day1.P2, result: Day1.R2) throws {
        try #expect(Day1.task2(data: data, param: param) == result)
    }
}

extension Day1 {
    static var example1 = """
3   4
4   3
2   5
1   3
3   9
3   3
"""
}

struct TestCase {
    
}
