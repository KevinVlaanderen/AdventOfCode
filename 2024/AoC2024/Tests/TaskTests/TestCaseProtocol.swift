import Testing
import Framework

protocol TestCaseProtocol: CustomTestStringConvertible, Sendable {
    var testDescription: String {get}
    
    func runTest() throws
}


extension CaseProtocol {
    func runTest() throws {
        try #expect(execute() == expected)
    }
}

extension Case: TestCaseProtocol {
    public var testDescription: String {
        return description
    }
}
