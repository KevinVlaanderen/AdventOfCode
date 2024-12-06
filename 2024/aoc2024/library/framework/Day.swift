protocol Day {
    associatedtype P = Void
    associatedtype R: Comparable
    
    func task1(data: String, param: P) async throws -> R
    func task2(data: String, param: P) async throws -> R
}
