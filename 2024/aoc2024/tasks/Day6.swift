import Foundation
internal import Algorithms
internal import SwiftGraph

struct Day6: Day {
    typealias R = Int

    func task1(data: String, param: P) throws -> R {
        let grid: any Grid<Content> = parse(data)

        guard let (guardPosition, guardDirection) = findGuard(grid: grid) else {
            fatalError("no guard found")
        }

        return PathTracker(grid: grid, startPosition: guardPosition, startDirection: guardDirection)
            .map { step in step.position }
            .uniqued()
            .count { _ in true }
    }
    
    func task2(data: String, param: P) async throws -> R {
        let grid: any Grid<Content> = parse(data)
        
        guard let (guardPosition, guardDirection) = findGuard(grid: grid) else {
            fatalError("no guard found")
        }

        return PathTracker(grid: grid, startPosition: guardPosition, startDirection: guardDirection)
            .filter(isValidTargetForObstruction(grid: grid, guardPosition: guardPosition))
            .uniqued(on: nextPosition)
            .count(where: entersLoop(grid: grid))
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
    
    private func findGuard(grid: any Grid<Content>) -> (Point, Direction)? {
        for item in grid {
            if case let .patrolGuard(guardDirection) = item.value {
                return (item.position, guardDirection)
            }
        }
        return nil
    }
    
    private func isValidTargetForObstruction(grid: any Grid<Content>, guardPosition: Point) -> (Step) -> Bool {
        return { step in
            let nextPosition = nextPosition(for: step)
            guard let content = grid[nextPosition], content != .obstruction else {
                return false
            }
            return nextPosition != guardPosition
        }
    }
    
    private func nextPosition(for step: Step) -> Point {
        step.position.neighbour(direction: step.direction)
    }
    
    private func entersLoop(grid: any Grid<Content>) -> (Step) -> Bool {
        return { step in
            var newGrid = grid
            newGrid[step.position.neighbour(direction: step.direction)] = .obstruction
            
            var seen: [Step] = []
            
            for step in PathTracker(grid: newGrid, startPosition: step.position, startDirection: step.direction) {
                let positionAhead = nextPosition(for: step)
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
