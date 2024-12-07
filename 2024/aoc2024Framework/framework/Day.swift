public protocol Day: Sendable {
    associatedtype P: Sendable = Void
    associatedtype R: Comparable, Sendable
    
    func perform(task: Task, data: String, param: P) throws -> R
}

public enum Task: Sendable {
    case task1, task2
}
