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
        case (.N, .CW90): .E
        case (.NE, .CW90): .SE
        case (.E, .CW90): .S
        case (.SE, .CW90): .SW
        case (.S, .CW90): .W
        case (.SW, .CW90): .NW
        case (.W, .CW90): .N
        case (.NW, .CW90): .NE
        case (.N, .CCW90): .W
        case (.NE, .CCW90): .NW
        case (.E, .CCW90): .N
        case (.SE, .CCW90): .NE
        case (.S, .CCW90): .E
        case (.SW, .CCW90): .SE
        case (.W, .CCW90): .S
        case (.NW, .CCW90): .SW
        case (.N, .Invert): .S
        case (.NE, .Invert): .SW
        case (.E, .Invert): .W
        case (.SE, .Invert): .NW
        case (.S, .Invert): .N
        case (.SW, .Invert): .NE
        case (.W, .Invert): .E
        case (.NW, .Invert): .SE
        }
    }
    
    func opposite() -> Direction {
        switch self {
        case .N: .S
        case .NE: .SW
        case .E: .W
        case .SE: .NW
        case .S: .N
        case .SW: .NE
        case .W: .E
        case .NW: .SE
        }
    }
     
    static var orthogonal: [Direction] { [.N, .E, .S, .W] }
}
