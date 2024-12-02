import Testing
@testable import aoc2024

final class Day2Tests {
    @Test(arguments: [
        (Day2.example1, (), 2),
        (Day2.data, (), 442)
    ])
    func task1(data: String, param: Day2.P1, result: Day2.R1) throws {
        try #expect(Day2.task1(data: data, param: param) == result)
    }
    
    @Test(arguments: [
        (Day2.example1, (), 4),
        (Day2.data, (), -1)
    ])
    func task2(data: String, param: Day2.P2, result: Day2.R2) throws {
        try #expect(Day2.task2(data: data, param: param) == result)
    }
}

extension Day2 {
    static var example1 = """
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
"""
}
