public struct DataDescriptor: Sendable {
    public init(url: String, load: @escaping DataLoader) {
        self.url = url
        self.load = load
    }
    
    public let url: String
    public let load: DataLoader
}

public typealias DataLoader = @Sendable () throws -> String

public func execute<D: Day>(day: D, task: Task, data: DataDescriptor, param: D.P) throws -> D.R {
    let _data = try data.load()
    
    return try day.perform(task: task, data: _data, param: param)
}
