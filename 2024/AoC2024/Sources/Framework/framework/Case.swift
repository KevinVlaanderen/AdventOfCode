public protocol CaseProtocol: Sendable {
    associatedtype D: Day
    
    var day: D {get}
    var task: Task {get}
    var expected: D.R {get}
    
    var description: String {get}

    func execute(data: Data) throws -> D.R
}

public struct Case<D: Day>: CaseProtocol {
    public let day: D
    public let task: Task
    public let expected: D.R
    
    let data: KeyPath<Data, String> & Sendable
    let param: D.P

    public init(day: D, task: Task, data: KeyPath<Data, String> & Sendable, param: D.P, expected: D.R) {
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
        var dataString = data.debugDescription
        if let index = dataString.firstIndex(of: ".") {
            dataString = String(dataString.suffix(from: index))
        }
        return "\(dayString) \(task) \(dataString)"
    }
    
    public func execute(data: Data) throws -> D.R {
        return try day.perform(task: task, data: data[keyPath: self.data], param: param)
    }
}
