import aoc2024Tasks
import aoc2024Framework

public func getCases() -> [any TestCaseProtocol] {
    return cases
}

public let cases: [any TestCaseProtocol] = [
    TestCase(day: Day1(), task: .task1, data: \.example1, param: (), expected: 11),
    TestCase(day: Day1(), task: .task1, data: \.data, param: (), expected: 1151792),
    TestCase(day: Day1(), task: .task2, data: \.example1, param: (), expected: 31),
    TestCase(day: Day1(), task: .task2, data: \.data, param: (), expected: 21790168),
    TestCase(day: Day2(), task: .task1, data: \.example1, param: (), expected: 2),
    TestCase(day: Day2(), task: .task1, data: \.data, param: (), expected: 442),
    TestCase(day: Day2(), task: .task2, data: \.example1, param: (), expected: 4),
    TestCase(day: Day2(), task: .task2, data: \.data, param: (), expected: 493),
    TestCase(day: Day3(), task: .task1, data: \.example1, param: (), expected: 161),
    TestCase(day: Day3(), task: .task1, data: \.data, param: (), expected: 180233229),
    TestCase(day: Day3(), task: .task2, data: \.example2, param: (), expected: 48),
    TestCase(day: Day3(), task: .task2, data: \.data, param: (), expected: 95411583),
    TestCase(day: Day4(), task: .task1, data: \.example1, param: "XMAS", expected: 18),
    TestCase(day: Day4(), task: .task1, data: \.data, param: "XMAS", expected: 2573),
    TestCase(day: Day4(), task: .task2, data: \.example1, param: "MAS", expected: 9),
    TestCase(day: Day4(), task: .task2, data: \.data, param: "MAS", expected: 1850),
    TestCase(day: Day5(), task: .task1, data: \.example1, param: (), expected: 143),
    TestCase(day: Day5(), task: .task1, data: \.data, param: (), expected: 4959),
    TestCase(day: Day5(), task: .task2, data: \.example1, param: (), expected: 123),
    TestCase(day: Day5(), task: .task2, data: \.data, param: (), expected: 4655),
    TestCase(day: Day6(), task: .task1, data: \.example1, param: (), expected: 41),
    TestCase(day: Day6(), task: .task1, data: \.data, param: (), expected: 4973),
    TestCase(day: Day6(), task: .task2, data: \.example1, param: (), expected: 6),
    TestCase(day: Day6(), task: .task2, data: \.data, param: (), expected: 1482),
    TestCase(day: Day7(), task: .task1, data: \.example1, param: [.add, .multiply], expected: 3749),
    TestCase(day: Day7(), task: .task1, data: \.data, param: [.add, .multiply], expected: 5702958180383),
    TestCase(day: Day7(), task: .task2, data: \.example1, param: Day7.Operators.allCases, expected: 11387),
    TestCase(day: Day7(), task: .task2, data: \.data, param: Day7.Operators.allCases, expected: 92612386119138)
]
