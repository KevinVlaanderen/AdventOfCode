import Foundation
internal import Algorithms
internal import SwiftGraph

struct Day6: Day {
    typealias R = Int

    func task1(data: String, param: P) throws -> R {
        let grid = parse(data)
        
        guard let guardItem = grid.first(where: { $0.value == .patrolGuard(.N) }) else {
            fatalError("no guard found")
        }

        return PathTracker(grid: grid, start: guardItem.position, direction: Direction.N)
            .map({ $0.position })
            .uniqued()
            .count(where: { _ in true })
    }
    
    func task2(data: String, param: P) throws -> R {
        let grid = parse(data)
        
        guard let guardItem = grid.first(where: { $0.value == .patrolGuard(.N) }) else {
            fatalError("no guard found")
        }
        
        return PathTracker(grid: grid, start: guardItem.position, direction: Direction.N)
            .map({ $0.position.neighbour(direction: $0.direction)})
            .filter({ $0 != guardItem.position && grid[$0] != nil })
            .uniqued()
            .count { newObstructionPosition in
                var newGrid = grid
                newGrid[newObstructionPosition] = .obstruction
                
                var seen: [Step] = []
                for step in PathTracker(grid: newGrid, start: guardItem.position, direction: Direction.N) {
                    let positionAhead = step.position.neighbour(direction: step.direction)
                    if newGrid[positionAhead] == .obstruction {
                        if seen.contains(step) {
                            return true
                        } else {
                            seen.append(step)
                        }
                    }
                }
                
                return false
            }
    }
    
    private func parse(_ data: String) -> ArrayGrid<Day6.Content> {
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
    
    struct Step: Hashable {
        let position: Point
        let direction: Direction
    }
    
    struct PathTracker: Sequence {
        let grid: ArrayGrid<Content>
        let start: Point
        let direction: Direction
        
        func makeIterator() -> PathIterator {
            return PathIterator(grid: grid, current: start, direction: direction)
        }
        
        struct PathIterator: IteratorProtocol {
            let grid: ArrayGrid<Content>
            var current: Point
            var direction: Direction
            var started = false
            
            mutating func next() -> Step? {
                if !started {
                    started = true
                    return Step(position: current, direction: direction)
                }

                let nextPosition = current.neighbour(direction: direction)

                switch grid[nextPosition] {
                case .obstruction:
                    direction = direction.rotate(.CW90)
                    return Step(position: current, direction: direction)
                case nil:
                    return nil
                default:
                    current = nextPosition
                    return Step(position: current, direction: direction)
                }
            }
        }
    }
}
