public enum Heading: CaseIterable, Codable {
    case north, northEast, east, southEast, south, southWest, west, northWest
}

public enum HeadingChange: CaseIterable {
    case clockwise90, counterClockwise90, invert
}

public extension Heading {
    func asPoint() -> Point {
        switch self {
        case .north:
            return Point(x: 0, y: -1)
        case .northEast:
            return Point(x: 1, y: -1)
        case .east:
            return Point(x: 1, y: 0)
        case .southEast:
            return Point(x: 1, y: 1)
        case .south:
            return Point(x: 0, y: 1)
        case .southWest:
            return Point(x: -1, y: 1)
        case .west:
            return Point(x: -1, y: 0)
        case .northWest:
            return Point(x: -1, y: -1)
        }
    }
    
    func rotate(_ change: HeadingChange) -> Heading {
        switch (self, change) {
        case (.north, .clockwise90): .east
        case (.northEast, .clockwise90): .southEast
        case (.east, .clockwise90): .south
        case (.southEast, .clockwise90): .southWest
        case (.south, .clockwise90): .west
        case (.southWest, .clockwise90): .northWest
        case (.west, .clockwise90): .north
        case (.northWest, .clockwise90): .northEast
        case (.north, .counterClockwise90): .west
        case (.northEast, .counterClockwise90): .northWest
        case (.east, .counterClockwise90): .north
        case (.southEast, .counterClockwise90): .northEast
        case (.south, .counterClockwise90): .east
        case (.southWest, .counterClockwise90): .southEast
        case (.west, .counterClockwise90): .south
        case (.northWest, .counterClockwise90): .southWest
        case (.north, .invert): .south
        case (.northEast, .invert): .southWest
        case (.east, .invert): .west
        case (.southEast, .invert): .northWest
        case (.south, .invert): .north
        case (.southWest, .invert): .northEast
        case (.west, .invert): .east
        case (.northWest, .invert): .southEast
        }
    }
    
    func opposite() -> Heading {
        switch self {
        case .north: .south
        case .northEast: .southWest
        case .east: .west
        case .southEast: .northWest
        case .south: .north
        case .southWest: .northEast
        case .west: .east
        case .northWest: .southEast
        }
    }
}
