import Testing
@testable import aoc2024

final class Day4Tests {
    @Test(arguments: [
        (Day4.example1, "XMAS", 18),
        (Day4.data, "XMAS", 2573)
    ])
    func task1(data: String, param: Day4.P1, result: Day4.R1) throws {
        try #expect(Day4.task1(data: data, param: param) == result)
    }
}

extension Day4 {
    static var example1 = """
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
"""
}
