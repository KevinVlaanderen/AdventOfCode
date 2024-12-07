import Testing
@testable import aoc2024Framework
@testable import aoc2024Data

struct Tests {
    @Test(arguments: cases)
    func run(testCase: any TestCaseProtocol) async throws {
        try testCase.runTest()
    }
}
