import Foundation
internal import Algorithms
internal import SwiftGraph

struct Day6: Day {
    typealias R = Int

    func task1(data: String, param: P) throws -> R {
        var grid = parse(data)
        var visited = ArrayGrid<Bool>(width: grid.width, height: grid.height, defaultValue: false)
        
        guard let patrolGuard = grid.first(where: { $0.value == .patrolGuard(.N) }) else {
            fatalError("no guard found")
        }
        grid[patrolGuard.position] = .empty
        
        var guardPosition: Point? = patrolGuard.position
        var guardDirection = Direction.N
        
        while let position = guardPosition {
            visited[position] = true
            
            let next = position.neighbour(direction: guardDirection)
            if grid[next] == .obstruction {
                guardDirection = guardDirection.rotate(.CW90)
            } else if grid[next] != nil {
                guardPosition = next
            } else {
                guardPosition = nil
            }
        }
        
        return visited.count(where: { $0.value })
    }
    
    func task2(data: String, param: P) throws -> R {
        return 0
    }
    
    private func parse(_ data: String) -> ArrayGrid<Content> {
        return ArrayGrid(data.split(whereSeparator: \.isNewline).map { line in
            Array(line).map({
                switch $0 {
                case ".":
                    return .empty
                case "#":
                    return .obstruction
                case "^":
                    return .patrolGuard(.N)
                default:
                    fatalError("unknown grid item: \($0)")
                }
            })
        })
    }
    
    enum Content: Equatable {
        case empty, obstruction
        case patrolGuard(Direction)
    }
}
