import Foundation
internal import Algorithms
internal import SwiftGraph
import aoc2024Framework

public struct Day6: Day {
    public typealias R = Int
    
    public init() {}
    
    public func perform(task: Task, data: String, param: P) throws -> Int {
        let grid: any Grid<Content> = parse(data)

        guard let (guardPosition, guardDirection) = findGuard(grid: grid) else {
            fatalError("no guard found")
        }
        
        switch task {
        case .task1:
            return try task1(grid: grid, startPosition: guardPosition, startDirection: guardDirection)
        case .task2:
            return try task2(grid: grid, startPosition: guardPosition, startDirection: guardDirection)
        @unknown default:
            fatalError("unknown task")
        }
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
    
    private func task1(grid: any Grid<Content>, startPosition: Point, startDirection: Direction) throws -> R {
        PathTracker(grid: grid, startPosition: startPosition, startDirection: startDirection)
            .map { step in step.position }
            .uniqued()
            .count { _ in true }
    }
    
    private func task2(grid: any Grid<Content>, startPosition: Point, startDirection: Direction) throws -> R {
        PathTracker(grid: grid, startPosition: startPosition, startDirection: startDirection)
            .filter(isValidTargetForObstruction(grid: grid, guardPosition: startPosition))
            .uniqued(on: nextPosition)
            .count(where: entersLoop(grid: grid))
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
