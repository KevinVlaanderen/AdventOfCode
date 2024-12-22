public protocol DayProtocol {
    associatedtype P
    associatedtype R: Comparable
    
    init(data: String, param: P)
    
    func perform() throws -> R
}

public extension DayProtocol {
    static func task(_ dataDescriptor: DataDescriptor, param: Self.P) -> Self? {
        do {
            let data = try dataDescriptor.load()
            return Self(data: data, param: param)
        } catch {
            return nil
        }
    }
}

open class Day<P, R>: DayProtocol where R: Comparable {
//    public typealias P = Task
//    public typealias R = Int
    
    public let data: String
    public let param: P
    
    public required init(data: String, param: P) {
        self.data = data
        self.param = param
    }
    
    open func perform() throws -> R {
        throw AoCError.notImplemented
    }
}

public enum Task {
    case task1, task2
}
