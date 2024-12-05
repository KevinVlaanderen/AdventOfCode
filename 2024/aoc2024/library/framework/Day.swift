protocol Day {
    associatedtype P = Void
    associatedtype R: Comparable
    
    func task1(data: String, param: P) throws -> R
    func task2(data: String, param: P) throws -> R
}
