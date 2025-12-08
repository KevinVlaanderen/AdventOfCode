public protocol DayProtocol {
    associatedtype P
    associatedtype R: Comparable
    
    init()
    func perform(data: String, task: Task, param: P) throws -> R
}

public extension DayProtocol {
    static func task(_ dataDescriptor: DataDescriptor, task: Task, param: Self.P) -> () throws -> R {
        {
            let data = try dataDescriptor.load()
            let day = Self()
            return try day.perform(data: data, task: task, param: param)
        }
    }
}

open class Day<P, R>: DayProtocol where R: Comparable {
    public required init() {}
    
    open func perform(data: String, task: Task, param: P) throws -> R {
        throw AoCError.notImplemented
    }
}

public enum Task {
    case task1, task2
}
