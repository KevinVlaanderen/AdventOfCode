import Foundation
internal import Algorithms
internal import HeapModule
import Framework

public struct Day20: Day {
    public typealias P = (maxCheats: Int, cutoff: Int)
    
    private let data: String
    private let param: P
    
    public init(data: String, param: P) {
        self.data = data
        self.param = param
    }
    
    public func perform() throws -> R {
        let grid = try parse(data: data)
        let track = try RaceTrack(from: grid)
        
        let path = try track.path(from: track.start, to: track.end)
                
        var cheats: [ShortcutKey: Int] = [:]
        
        for step in path.enumerated() {
            let stepsWithinDistance = path.enumerated().filter({ $0.offset > step.offset+1})
                .map({ ($0, step.element.manhattanDistance(to: $0.element)) })
                .filter({ $0.1 <= param.maxCheats })
            
            for (otherStep, distance) in stepsWithinDistance where otherStep.offset > step.offset + distance {
                let key = ShortcutKey(from: step.element, to: otherStep.element)
                cheats[key] = otherStep.offset - step.offset - distance
            }
        }
                
        return cheats.filter({ $0.value >= param.cutoff }).count
    }
    
    private func parse(data: String) throws -> any Grid<TileType> {
        try ArrayGrid(data.lines.map { line throws in
            try line.map { character throws in
                switch character {
                case ".": .empty
                case "#": .wall
                case "S": .start
                case "E": .end
                default: throw AoCError.parseError("Unknown character: \(character)")
                }
            }
        })
    }
}

private struct ShortcutKey: Hashable {
    let from: Point
    let to: Point
}

private struct RaceTrack {
    let start: Point
    let end: Point
    var grid: any Grid<TileType>
    
    init(from grid: any Grid<TileType>) throws {
        self.grid = grid
        
        guard let start = grid.first(where: { $0.value == TileType.start }) else {
            throw AoCError.parseError("No start position")
        }
        guard let end = grid.first(where: { $0.value == TileType.end }) else {
            throw AoCError.parseError("No end position")
        }
        
        self.start = start.position
        self.end = end.position
    }
    
    func path(from start: Point, to end: Point) throws -> [Point] {
        var path = [start]
        
        var current = start
        var previous: Point? = nil
        
        while current != end {
            for neighbour in current.neighbours(directions: Direction.allCases) where neighbour != previous && grid[neighbour] != .wall {
                path.append(neighbour)
                previous = current
                current = neighbour
                break
            }
        }
        
        return path
    }
}

private enum TileType: Equatable, Hashable {
    case empty, wall, start, end
}
