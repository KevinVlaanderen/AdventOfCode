public struct Point: Hashable, Sendable, Comparable {
    public var x, y: Int

    public init(x: Int, y: Int) {
        self.x = x
        self.y = y
    }
    
    public func neighbour(heading: Heading, distance: Int = 1) -> Point {
        var neighbour = self
        let offset = heading.asPoint()
        neighbour.x += offset.x * distance
        neighbour.y += offset.y * distance
        return neighbour
    }
    
    public func neighbour(direction: Direction, distance: Int = 1) -> Point {
        var neighbour = self
        let offset = direction.asPoint()
        neighbour.x += offset.x * distance
        neighbour.y += offset.y * distance
        return neighbour
    }
    
    public func neighbours(headings: [Heading]) -> [Point] {
        headings.map { self.neighbour(heading: $0) }
    }
    
    public func neighbours(directions: [Direction]) -> [Point] {
        directions.map { self.neighbour(direction: $0) }
    }
    
    public static func < (lhs: Point, rhs: Point) -> Bool {
        lhs.x == rhs.x && lhs.y == rhs.y
    }
    
    public func manhattanDistance(to point: Point) -> Int {
        abs(self.x-point.x)+abs(self.y-point.y)
    }
}
