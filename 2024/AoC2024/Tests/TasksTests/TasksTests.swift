import Testing
@testable import Framework
@testable import Data
@testable import Tasks

struct Tests {
    let data = Data()
    
    @Test(arguments: cases)
    func run(c: any CaseProtocol) async throws {
        try c.runTest(data: data)
    }
    
    @Test()
    func runSpecific() async throws {
        let c = cases.filter { $0.day is Day6 && $0.task == .task2 }[1]
        try c.runTest(data: data)
    }
}
