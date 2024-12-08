import Testing
@testable import Framework
@testable import Data
@testable import Tasks

struct Tests {
    @Test(arguments: cases)
    func run(c: any CaseProtocol) async throws {
        try c.runTest()
    }
    
    @Test()
    func runSpecific() async throws {
        let c = cases.filter { $0.day is Day8 && $0.task == .task1 }[0]
        try c.runTest()
    }
}
