import Testing
import Framework

protocol TestCaseProtocol: CustomTestStringConvertible, Sendable {
    var testDescription: String {get}
    
    func runTest(data: Data) throws
}


extension CaseProtocol {
    func runTest(data: Data) throws {
        try #expect(execute(data: data) == expected)
    }
}

extension Case: TestCaseProtocol {
    public var testDescription: String {
        return description
    }
}
