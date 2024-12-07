import Testing

public protocol TestSuite {
    associatedtype D: Day
    
    static var cases: [TestCase<D>] {get}
    
    func run(testCase: TestCase<D>) async throws
}

public protocol TestCaseProtocol: CustomTestStringConvertible, Sendable {
    associatedtype D: Day
    
    var testDescription: String {get}
    
    func execute() throws
    func runTest() throws
}

public struct TestCase<D: Day>: TestCaseProtocol {
    let day: D
    let task: Task
    let data: KeyPath<D, String> & Sendable
    let param: D.P
    let expected: D.R

    public init(day: D, task: Task, data: KeyPath<D, String> & Sendable, param: D.P, expected: D.R) {
        self.day = day
        self.task = task
        self.data = data
        self.param = param
        self.expected = expected
    }
    
    public var testDescription: String {
        return "\(task) \(data)"
    }
    
    public func execute() throws {
        _ = try day.perform(task: task, data: day[keyPath: data], param: param)
    }
    
    public func runTest() throws {
        try #expect(day.perform(task: task, data: day[keyPath: data], param: param) == expected)
    }
}
