import Foundation
internal import Algorithms
internal import SwiftGraph

struct Day6: Day {
    typealias P = Direction
    typealias R = Int

    func task1(data: String, param: P) throws -> R {
        let grid: any Grid<Content> = parse(data)

        guard let guardItem = grid.first(where: { $0.value == .patrolGuard(param) }) else {
            fatalError("no guard found")
        }

        return PathTracker(grid: grid, startPosition: guardItem.position, startDirection: param)
            .map { step in step.position }
            .uniqued()
            .count { _ in true }
    }
    
    func task2(data: String, param: P) throws -> R {
        var grid: any Grid<Content> = parse(data)
        
        guard let guardItem = grid.first(where: { $0.value == .patrolGuard(param) }) else {
            fatalError("no guard found")
        }

        return PathTracker(grid: grid, startPosition: guardItem.position, startDirection: param)
            .map { step in step.position.neighbour(direction: step.direction) }
            .filter { position in
                guard let content = grid[position], content != .obstruction else {
                    return false
                }
                return position != guardItem.position
            }
            .uniqued()
            .count { position in checkForLoop(grid: &grid, obstructionPosition: position, startPosition: guardItem.position, startDirection: param) }
    }
    
    private func parse(_ data: String) -> some Grid<Content> {
        ArrayGrid(data.split(whereSeparator: \.isNewline).map { line in
            Array(line).map { charater in
                switch charater {
                case ".":
                    return .empty
                case "#":
                    return .obstruction
                case "^":
                    return .patrolGuard(.N)
                default:
                    fatalError("unknown grid item: \(charater)")
                }
            }
        })
    }
    
    private func checkForLoop(grid: inout any Grid<Content>, obstructionPosition: Point, startPosition: Point, startDirection: Direction) -> Bool {
        if grid[obstructionPosition] == .obstruction {
            return false
        }
        
        grid[obstructionPosition] = .obstruction
        
        var seen: [Step] = []
        var result = false
        
        for step in PathTracker(grid: grid, startPosition: startPosition, startDirection: startDirection) {
            let positionAhead = step.position.neighbour(direction: step.direction)
            if grid[positionAhead] == .obstruction {
                if seen.contains(step) {
                    result = true
                    break
                } else {
                    seen.append(step)
                }
            }
        }
        
        grid[obstructionPosition] = .empty
        return result
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
        let grid: any Grid<Content>
        let startPosition: Point
        let startDirection: Direction
        
        func makeIterator() -> PathIterator {
            return PathIterator(grid: grid, currentPosition: startPosition, currentDirection: startDirection)
        }
        
        struct PathIterator: IteratorProtocol {
            let grid: any Grid<Content>
            var currentPosition: Point
            var currentDirection: Direction
            var started = false
            
            mutating func next() -> Step? {
                if !started {
                    started = true
                    return Step(position: currentPosition, direction: currentDirection)
                }

                let nextPosition = currentPosition.neighbour(direction: currentDirection)

                switch grid[nextPosition] {
                case .obstruction:
                    currentDirection = currentDirection.rotate(.CW90)
                    return Step(position: currentPosition, direction: currentDirection)
                case nil:
                    return nil
                default:
                    currentPosition = nextPosition
                    return Step(position: currentPosition, direction: currentDirection)
                }
            }
        }
    }
}
