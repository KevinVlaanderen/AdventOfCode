import Foundation
internal import Algorithms
internal import HeapModule
import Framework

public final class Day20: Day<(maxCheats: Int, cutoff: Int), Int> {
    public override func perform(data: String, task: Task, param: P) throws -> R {
        let grid = try parse(data)
        let track = try RaceTrack(from: grid)
        
        let path = try track.path(from: track.start, to: track.end)

        return path.enumerated().reduce(into: 0) { result, step in
            let stepDistance = step.offset
            
            for otherStep in path[stepDistance+1..<path.endIndex].enumerated() {
                let otherStepDistance = stepDistance + otherStep.offset + 1
                let distanceBetween = step.element.manhattanDistance(to: otherStep.element)

                if distanceBetween <= param.maxCheats && otherStepDistance > step.offset + distanceBetween {
                    let cheatDistance = otherStepDistance - step.offset - distanceBetween

                    if cheatDistance >= param.cutoff {
                        result += 1
                    }
                }
            }
        }
    }
    
    private func parse(_ data: String) throws -> any Grid<TileType> {
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
