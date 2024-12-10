import Foundation
import Framework
internal import Algorithms

public struct Day10: Day {
    public init() {}
        
    public func perform(task: Task, data: String, param: P) throws -> Int {
        let topographicMap = parse(data)
        
        let trailheads = topographicMap.filter {
            if case let .passable(height) = $0.value, height == 0 {
                return true
            }
            return false
        }
        
        return trailheads.reduce(0) { result, current in
            var found: [Point] = []
            return result + pathsTo(9, from: current.position, topographicMap: topographicMap, found: &found)
        }
    }
    
    private func parse(_ data: String) -> any Grid<Position> {
        ArrayGrid(data.split(whereSeparator: \.isNewline).map { line in
            Array(line).map { character in
                switch character {
                case ".":
                    return .impassable
                default:
                    return .passable(character.wholeNumberValue!)
                }
            }
        })
    }
    
    private func pathsTo(_ target: Int, from: Point, topographicMap: any Grid<Position>, found: inout [Point]) -> Int {
        let newPosition = topographicMap[from]
        
        guard case let .passable(height) = newPosition else {
            return 0
        }
        
        if height == target {
            if found.contains(from) {
                return 0
            } else {
                found.append(from)
                return 1
            }
        }
        
        let directions: [Direction] = [.N, .E, .S, .W]
        
        return directions.reduce(0) { result, direction in
            let next = from.neighbour(direction: direction)
            let nextPosition = topographicMap[next]
            guard case let .passable(nextHeight) = nextPosition, nextHeight == height+1 else {
                return result
            }
            return result + pathsTo(target, from: next, topographicMap: topographicMap, found: &found)
        }
    }
    
    private enum Position: Equatable {
        case passable(Int)
        case impassable
    }
}

