public enum Direction: CaseIterable, Codable {
    case left, right, up, down
}

public enum Rotation: CaseIterable {
    case left, right, invert
}

public extension Direction {
    func asPoint() -> Point {
        switch self {
        case .left:             Point(x: -1, y: 0)
        case .right:            Point(x: 1, y: 0)
        case .up:               Point(x: 0, y: -1)
        case .down:             Point(x: 0, y: 1)
        }
    }
    
    func rotate(_ change: Rotation) -> Direction {
        switch (self, change) {
        case (.left, .left):    .down
        case (.left, .right):   .up
        case (.left, .invert):  .right
        case (.right, .left):   .up
        case (.right, .right):  .down
        case (.right, .invert): .left
        case (.up, .left):      .left
        case (.up, .right):     .right
        case (.up, .invert):    .down
        case (.down, .left):    .right
        case (.down, .right):   .left
        case (.down, .invert):  .up
        }
    }
    
    func opposite() -> Direction {
        switch self {
        case .left:             .right
        case .right:            .left
        case .up:               .down
        case .down:             .up
        }
    }
}
