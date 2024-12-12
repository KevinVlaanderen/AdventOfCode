public enum Direction: CaseIterable {
    case N, NE, E, SE, S, SW, W, NW
}

public enum DirectionChange: CaseIterable {
    case CW90, CCW90, Invert
}

public extension Direction {
    func asPoint() -> Point {
        switch self {
        case .N:
            return Point(x: 0, y: -1)
        case .NE:
            return Point(x: 1, y: -1)
        case .E:
            return Point(x: 1, y: 0)
        case .SE:
            return Point(x: 1, y: 1)
        case .S:
            return Point(x: 0, y: 1)
        case .SW:
            return Point(x: -1, y: 1)
        case .W:
            return Point(x: -1, y: 0)
        case .NW:
            return Point(x: -1, y: -1)
        }
    }
    
    func rotate(_ change: DirectionChange) -> Direction {
        switch (self, change) {
        case (.N, .CW90):
            return .E
        case (.NE, .CW90):
            return .SE
        case (.E, .CW90):
            return .S
        case (.SE, .CW90):
            return .SW
        case (.S, .CW90):
            return .W
        case (.SW, .CW90):
            return .NW
        case (.W, .CW90):
            return .N
        case (.NW, .CW90):
            return .NE
        case (.N, .CCW90):
            return .W
        case (.NE, .CCW90):
            return .NW
        case (.E, .CCW90):
            return .N
        case (.SE, .CCW90):
            return .NE
        case (.S, .CCW90):
            return .E
        case (.SW, .CCW90):
            return .SE
        case (.W, .CCW90):
            return .S
        case (.NW, .CCW90):
            return .SW
        case (.N, .Invert):
            return .S
        case (.NE, .Invert):
            return .SW
        case (.E, .Invert):
            return .W
        case (.SE, .Invert):
            return .NW
        case (.S, .Invert):
            return .N
        case (.SW, .Invert):
            return .NE
        case (.W, .Invert):
            return .E
        case (.NW, .Invert):
            return .SE
        }
    }
     
    static var orthogonal: [Direction] { [.N, .E, .S, .W] }
}
