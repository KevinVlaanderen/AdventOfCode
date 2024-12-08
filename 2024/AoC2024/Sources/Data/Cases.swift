import Tasks
import Framework

public func getCases() -> [any CaseProtocol] {
    return cases
}

public let cases: [any CaseProtocol] = [
    Case(day: Day1(), task: .task1, data: \.example1, param: (), expected: 11),
    Case(day: Day1(), task: .task1, data: \.data, param: (), expected: 1151792),
    Case(day: Day1(), task: .task2, data: \.example1, param: (), expected: 31),
    Case(day: Day1(), task: .task2, data: \.data, param: (), expected: 21790168),
    Case(day: Day2(), task: .task1, data: \.example1, param: (), expected: 2),
    Case(day: Day2(), task: .task1, data: \.data, param: (), expected: 442),
    Case(day: Day2(), task: .task2, data: \.example1, param: (), expected: 4),
    Case(day: Day2(), task: .task2, data: \.data, param: (), expected: 493),
    Case(day: Day3(), task: .task1, data: \.example1, param: (), expected: 161),
    Case(day: Day3(), task: .task1, data: \.data, param: (), expected: 180233229),
    Case(day: Day3(), task: .task2, data: \.example2, param: (), expected: 48),
    Case(day: Day3(), task: .task2, data: \.data, param: (), expected: 95411583),
    Case(day: Day4(), task: .task1, data: \.example1, param: "XMAS", expected: 18),
    Case(day: Day4(), task: .task1, data: \.data, param: "XMAS", expected: 2573),
    Case(day: Day4(), task: .task2, data: \.example1, param: "MAS", expected: 9),
    Case(day: Day4(), task: .task2, data: \.data, param: "MAS", expected: 1850),
    Case(day: Day5(), task: .task1, data: \.example1, param: (), expected: 143),
    Case(day: Day5(), task: .task1, data: \.data, param: (), expected: 4959),
    Case(day: Day5(), task: .task2, data: \.example1, param: (), expected: 123),
    Case(day: Day5(), task: .task2, data: \.data, param: (), expected: 4655),
    Case(day: Day6(), task: .task1, data: \.example1, param: (), expected: 41),
    Case(day: Day6(), task: .task1, data: \.data, param: (), expected: 4973),
    Case(day: Day6(), task: .task2, data: \.example1, param: (), expected: 6),
    Case(day: Day6(), task: .task2, data: \.data, param: (), expected: 1482),
    Case(day: Day7(), task: .task1, data: \.example1, param: [.add, .multiply], expected: 3749),
    Case(day: Day7(), task: .task1, data: \.data, param: [.add, .multiply], expected: 5702958180383),
    Case(day: Day7(), task: .task2, data: \.example1, param: Day7.Operators.allCases, expected: 11387),
    Case(day: Day7(), task: .task2, data: \.data, param: Day7.Operators.allCases, expected: 92612386119138),
    Case(day: Day8(), task: .task1, data: \.example1, param: (), expected: 14),
    Case(day: Day8(), task: .task1, data: \.data, param: (), expected: 320)
]
