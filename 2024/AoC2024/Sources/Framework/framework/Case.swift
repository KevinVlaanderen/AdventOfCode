public protocol CaseProtocol: Sendable {
    associatedtype D: Day
    
    var day: D {get}
    var task: Task {get}
    var expected: D.R {get}
    
    var description: String {get}

    func execute() throws -> D.R
}

public struct DataDescriptor: Sendable {
    public init(url: String, load: @escaping DataLoader) {
        self.url = url
        self.load = load
    }
    
    public let url: String
    public let load: DataLoader
}

public typealias DataLoader = @Sendable () throws -> String

public struct Case<D: Day>: CaseProtocol, Sendable {
    public let day: D
    public let task: Task
    public let expected: D.R
    
    private let data: DataDescriptor
    let param: D.P

    public init(day: D, task: Task, data: DataDescriptor, param: D.P, expected: D.R) {
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
        var dataString = self.data.url
        if let slashIndex = dataString.lastIndex(of: "/"),
           let endIndex = dataString.lastIndex(of: ".") {
            let startIndex = dataString.index(after: slashIndex)
            dataString = String(dataString[startIndex..<endIndex])
        }
        return "\(dayString) \(task) \(dataString)"
    }
    
    public func execute() throws -> D.R {
        let data = try self.data.load()
        
        return try day.perform(task: task, data: data, param: param)
    }
}
