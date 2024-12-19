public protocol Day: Sendable {
    associatedtype P: Sendable = Task
    associatedtype R: Comparable, Sendable = Int
    
    init(data: String, param: P)
    
    func perform() throws -> R
}

public extension Day {
    static func task(_ dataDescriptor: DataDescriptor, param: Self.P) -> Self? {
        do {
            let data = try dataDescriptor.load()
            return Self(data: data, param: param)
        } catch {
            return nil
        }
    }
}

public enum Task: Sendable {
    case task1, task2
}
