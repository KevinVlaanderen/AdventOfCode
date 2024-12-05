import Testing
@testable import aoc2024

protocol TestSuite {
    associatedtype D: Day
    
    func task1(testCase: TestCase<D>) throws
    func task2(testCase: TestCase<D>) throws
}

struct TestCase<D: Day>: CustomTestStringConvertible {
    let day: D
    let data: KeyPath<D, String>
    let param: D.P
    let expected: D.R

    var testDescription: String {
        return "\(data)"
    }
    
    func execute(_ method: (String, D.P) throws -> D.R) throws {
        try #expect(method(day[keyPath: data], param) == expected)
    }
}
