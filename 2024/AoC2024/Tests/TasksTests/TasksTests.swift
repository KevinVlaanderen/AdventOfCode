import Testing
@testable import Framework
@testable import Data

struct Tests {
    @Test(arguments: cases)
    func run(c: any CaseProtocol) async throws {
        try c.runTest()
    }
}
