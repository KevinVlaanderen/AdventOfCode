import Testing
@testable import aoc2024

final class Day3Tests {
    @Test(arguments: [
        (Day3.example1, (), 161),
        (Day3.data, (), -1)
    ])
    func task1(data: String, param: Day3.P1, result: Day3.R1) throws {
        try #expect(Day3.task1(data: data, param: param) == result)
    }
    
    @Test(arguments: [
        (Day3.example1, (), -1),
        (Day3.data, (), -1)
    ])
    func task2(data: String, param: Day3.P2, result: Day3.R2) throws {
        try #expect(Day3.task2(data: data, param: param) == result)
    }
}

extension Day3 {
    static var example1 = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
}
