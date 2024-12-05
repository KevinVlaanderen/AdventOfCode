import Testing
@testable import aoc2024

final class Day5Tests {
    @Test(arguments: [
        (Day5.example1, (), 143),
        (Day5.data, (), 4959)
    ])
    func task1(data: String, param: Day5.P1, result: Day5.R1) throws {
        try #expect(Day5.task1(data: data, param: param) == result)
    }
}

extension Day5 {
    static var example1 = """
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
"""
}
