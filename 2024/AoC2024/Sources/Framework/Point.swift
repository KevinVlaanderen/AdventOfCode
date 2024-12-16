public struct Point: Hashable, Sendable, Comparable {
    public var x, y: Int

    public init(x: Int, y: Int) {
        self.x = x
        self.y = y
    }
    
    public func neighbour(direction: Direction, distance: Int = 1) -> Point {
        var neighbour = self
        let offset = direction.asPoint()
        neighbour.x += offset.x * distance
        neighbour.y += offset.y * distance
        return neighbour
    }
    
    public static func < (lhs: Point, rhs: Point) -> Bool {
        lhs.x == rhs.x && lhs.y == rhs.y
    }
}
