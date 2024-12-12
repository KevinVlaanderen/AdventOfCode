import Testing
@testable import Framework
@testable import Data
@testable import Tasks

struct Tests {
    @Test(arguments: await cases)
    func run(c: any CaseProtocol) async throws {
        try c.runTest()
    }
}
