import Foundation
import Tasks
import Framework

extension File: @unchecked Sendable {}

public func loadData(file: File) -> DataDescriptor {
    DataDescriptor(url: file.path, load: {
        try String(contentsOfFile: file.path).trimmingCharacters(in: .newlines)
    })
}

@MainActor
public let cases: [any CaseProtocol] = [
    Case(day: Day1(), task: .task1, data: loadData(file: Data.Days.Day1.example1), param: (), expected: 11),
    Case(day: Day1(), task: .task1, data: loadData(file: Data.Days.Day1.data), param: (), expected: 1151792),
    Case(day: Day1(), task: .task2, data: loadData(file: Data.Days.Day1.example1), param: (), expected: 31),
    Case(day: Day1(), task: .task2, data: loadData(file: Data.Days.Day1.data), param: (), expected: 21790168),
    Case(day: Day2(), task: .task1, data: loadData(file: Data.Days.Day2.example1), param: (), expected: 2),
    Case(day: Day2(), task: .task1, data: loadData(file: Data.Days.Day2.data), param: (), expected: 442),
    Case(day: Day2(), task: .task2, data: loadData(file: Data.Days.Day2.example1), param: (), expected: 4),
    Case(day: Day2(), task: .task2, data: loadData(file: Data.Days.Day2.data), param: (), expected: 493),
    Case(day: Day3(), task: .task1, data: loadData(file: Data.Days.Day3.example1), param: (), expected: 161),
    Case(day: Day3(), task: .task1, data: loadData(file: Data.Days.Day3.data), param: (), expected: 180233229),
    Case(day: Day3(), task: .task2, data: loadData(file: Data.Days.Day3.example2), param: (), expected: 48),
    Case(day: Day3(), task: .task2, data: loadData(file: Data.Days.Day3.data), param: (), expected: 95411583),
    Case(day: Day4(), task: .task1, data: loadData(file: Data.Days.Day4.example1), param: "XMAS", expected: 18),
    Case(day: Day4(), task: .task1, data: loadData(file: Data.Days.Day4.data), param: "XMAS", expected: 2573),
    Case(day: Day4(), task: .task2, data: loadData(file: Data.Days.Day4.example1), param: "MAS", expected: 9),
    Case(day: Day4(), task: .task2, data: loadData(file: Data.Days.Day4.data), param: "MAS", expected: 1850),
    Case(day: Day5(), task: .task1, data: loadData(file: Data.Days.Day5.example1), param: (), expected: 143),
    Case(day: Day5(), task: .task1, data: loadData(file: Data.Days.Day5.data), param: (), expected: 4959),
    Case(day: Day5(), task: .task2, data: loadData(file: Data.Days.Day5.example1), param: (), expected: 123),
    Case(day: Day5(), task: .task2, data: loadData(file: Data.Days.Day5.data), param: (), expected: 4655),
    Case(day: Day6(), task: .task1, data: loadData(file: Data.Days.Day6.example1), param: (), expected: 41),
    Case(day: Day6(), task: .task1, data: loadData(file: Data.Days.Day6.data), param: (), expected: 4973),
    Case(day: Day6(), task: .task2, data: loadData(file: Data.Days.Day6.example1), param: (), expected: 6),
    Case(day: Day6(), task: .task2, data: loadData(file: Data.Days.Day6.data), param: (), expected: 1482),
    Case(day: Day7(), task: .task1, data: loadData(file: Data.Days.Day7.example1), param: [.add, .multiply], expected: 3749),
    Case(day: Day7(), task: .task1, data: loadData(file: Data.Days.Day7.data), param: [.add, .multiply], expected: 5702958180383),
    Case(day: Day7(), task: .task2, data: loadData(file: Data.Days.Day7.example1), param: Day7.Operators.allCases, expected: 11387),
    Case(day: Day7(), task: .task2, data: loadData(file: Data.Days.Day7.data), param: Day7.Operators.allCases, expected: 92612386119138),
    Case(day: Day8(), task: .task1, data: loadData(file: Data.Days.Day8.example1), param: (), expected: 14),
    Case(day: Day8(), task: .task1, data: loadData(file: Data.Days.Day8.data), param: (), expected: 320),
    Case(day: Day8(), task: .task2, data: loadData(file: Data.Days.Day8.example1), param: (), expected: 34),
    Case(day: Day8(), task: .task2, data: loadData(file: Data.Days.Day8.data), param: (), expected: 1157),
    Case(day: Day9(), task: .task1, data: loadData(file: Data.Days.Day9.example1), param: (), expected: 1928),
    Case(day: Day9(), task: .task1, data: loadData(file: Data.Days.Day9.data), param: (), expected: 6356833654075),
    Case(day: Day9(), task: .task2, data: loadData(file: Data.Days.Day9.example1), param: (), expected: 2858),
    Case(day: Day9(), task: .task2, data: loadData(file: Data.Days.Day9.data), param: (), expected: 6389911791746),
    Case(day: Day10(), task: .task1, data: loadData(file: Data.Days.Day10.example1), param: (), expected: 1),
    Case(day: Day10(), task: .task1, data: loadData(file: Data.Days.Day10.example2), param: (), expected: 2),
    Case(day: Day10(), task: .task1, data: loadData(file: Data.Days.Day10.example3), param: (), expected: 4),
    Case(day: Day10(), task: .task1, data: loadData(file: Data.Days.Day10.example4), param: (), expected: 3),
    Case(day: Day10(), task: .task1, data: loadData(file: Data.Days.Day10.example5), param: (), expected: 36),
    Case(day: Day10(), task: .task1, data: loadData(file: Data.Days.Day10.data), param: (), expected: 816),
    Case(day: Day10(), task: .task2, data: loadData(file: Data.Days.Day10.example6), param: (), expected: 3),
    Case(day: Day10(), task: .task2, data: loadData(file: Data.Days.Day10.example7), param: (), expected: 13),
    Case(day: Day10(), task: .task2, data: loadData(file: Data.Days.Day10.example8), param: (), expected: 227),
    Case(day: Day10(), task: .task2, data: loadData(file: Data.Days.Day10.example5), param: (), expected: 81),
    Case(day: Day10(), task: .task2, data: loadData(file: Data.Days.Day10.data), param: (), expected: 1960),
    Case(day: Day11(), task: .task1, data: loadData(file: Data.Days.Day11.example1), param: 25, expected: 55312),
    Case(day: Day11(), task: .task1, data: loadData(file: Data.Days.Day11.data), param: 25, expected: 213625),
    Case(day: Day11(), task: .task2, data: loadData(file: Data.Days.Day11.data), param: 75, expected: 252442982856820),
    Case(day: Day12(), task: .task1, data: loadData(file: Data.Days.Day12.example1), param: (), expected: 140),
    Case(day: Day12(), task: .task1, data: loadData(file: Data.Days.Day12.example2), param: (), expected: 772),
    Case(day: Day12(), task: .task1, data: loadData(file: Data.Days.Day12.example3), param: (), expected: 1930),
    Case(day: Day12(), task: .task1, data: loadData(file: Data.Days.Day12.data), param: (), expected: 1457298),
    Case(day: Day12(), task: .task2, data: loadData(file: Data.Days.Day12.example1), param: (), expected: 80),
    Case(day: Day12(), task: .task2, data: loadData(file: Data.Days.Day12.example2), param: (), expected: 436),
    Case(day: Day12(), task: .task2, data: loadData(file: Data.Days.Day12.example4), param: (), expected: 236),
    Case(day: Day12(), task: .task2, data: loadData(file: Data.Days.Day12.example5), param: (), expected: 368),
    Case(day: Day12(), task: .task2, data: loadData(file: Data.Days.Day12.data), param: (), expected: 921636),
]
