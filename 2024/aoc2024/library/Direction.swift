enum Direction: CaseIterable {
    case N, NE, E, SE, S, SW, W, NW
}

extension Direction {
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
}
