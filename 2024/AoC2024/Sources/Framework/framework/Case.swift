public protocol CaseProtocol: Sendable {
    associatedtype D: Day
    
    var day: D {get}
    var task: Task {get}
    var data: KeyPath<D, String> & Sendable {get}
    var param: D.P {get}
    var expected: D.R {get}
    
    var description: String {get}

    func execute() throws -> D.R
}

public struct Case<D: Day>: CaseProtocol {
    public let day: D
    public let task: Task
    public let data: KeyPath<D, String> & Sendable
    public let param: D.P
    public let expected: D.R

    public init(day: D, task: Task, data: KeyPath<D, String> & Sendable, param: D.P, expected: D.R) {
        self.day = day
        self.task = task
        self.data = data
        self.param = param
        self.expected = expected
    }
    
    public var description: String {
        var dayString = "\(day)"
        if let index = dayString.firstIndex(of: "(") {
            dayString = String(dayString.prefix(upTo: index))
        }
        return "\(dayString).\(task).\(data.debugDescription.split(separator: ".")[1])"
    }
    
    public func execute() throws -> D.R {
        return try day.perform(task: task, data: day[keyPath: data], param: param)
    }
}
