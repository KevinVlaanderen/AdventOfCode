protocol Day {
    associatedtype P1 = Void
    associatedtype P2 = Void
    associatedtype R1: Comparable
    associatedtype R2: Comparable
    
    static func task1(data: String, param: P1) throws -> R1
    static func task2(data: String, param: P2) throws -> R2
}
