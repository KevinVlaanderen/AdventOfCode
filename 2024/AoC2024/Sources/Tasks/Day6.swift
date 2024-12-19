import Foundation
internal import Algorithms
internal import SwiftGraph
import Framework

public struct Day6: Day {
    private let data: String
    private let param: P
    
    public init(data: String, param: P) {
        self.data = data
        self.param = param
    }
    
    public func perform() throws -> Int {
        let grid: any Grid<Content> = parse(data)

        guard let (guardPosition, guardDirection) = findGuard(grid: grid) else {
            fatalError("no guard found")
        }
        
        return switch param {
        case .task1: try task1(grid: grid, startPosition: guardPosition, startDirection: guardDirection)
        case .task2: try task2(grid: grid, startPosition: guardPosition, startDirection: guardDirection)
        }
    }
    
    private func parse(_ data: String) -> some Grid<Content> {
        ArrayGrid(data.lines.map { line in
            line.map { character in
                switch character {
                case ".":   .empty
                case "#":   .obstruction
                case "^":   .patrolGuard(.N)
                default:    fatalError("unknown grid item: \(character)")
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
    
    private func task1(grid: any Grid<Content>, startPosition: Point, startDirection: Direction) throws -> Int {
        PathTracker(grid: grid, startPosition: startPosition, startDirection: startDirection)
            .map { step in step.position }
            .uniqued()
            .count { _ in true }
    }
    
    private func task2(grid: any Grid<Content>, startPosition: Point, startDirection: Direction) throws -> Int {
        PathTracker(grid: grid, startPosition: startPosition, startDirection: startDirection)
            .filter(isValidTargetForObstruction(grid: grid, guardPosition: startPosition))
            .uniqued(on: nextPosition)
            .count(where: entersLoop(grid: grid))
    }
    
    private func isValidTargetForObstruction(grid: any Grid<Content>, guardPosition: Point) -> (Step) -> Bool {
        { step in
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
        { outerStep in
            var newGrid = grid
            newGrid[outerStep.position.neighbour(direction: outerStep.direction)] = .obstruction
            
            var seen: [Step] = []
            
            for (step, nextStep) in PathTracker(grid: newGrid, startPosition: outerStep.position, startDirection: outerStep.direction).adjacentPairs() {
                if nextStep.content == .obstruction {
                    if seen.contains(step) {
                        return true
                    }
                    seen.append(step)
                }
            }
            
            return false
        }
    }
    
    enum Content: Equatable, Hashable {
        case empty, obstruction
        case patrolGuard(Direction)
    }
    
    struct Step: Hashable {
        let position: Point
        let direction: Direction
        let content: Content
    }
    
    struct PathTracker: Sequence {
        let grid: any Grid<Content>
        let startPosition: Point
        let startDirection: Direction
        
        func makeIterator() -> PathIterator {
            PathIterator(grid: grid, currentPosition: startPosition, currentDirection: startDirection)
        }
        
        struct PathIterator: IteratorProtocol {
            let grid: any Grid<Content>
            var currentPosition: Point
            var currentDirection: Direction
            var started = false
            
            mutating func next() -> Step? {
                if !started {
                    guard let currentContent = grid[currentPosition] else {
                        return nil
                    }
                    started = true
                    return Step(position: currentPosition, direction: currentDirection, content: currentContent)
                }

                let nextPosition = currentPosition.neighbour(direction: currentDirection)

                guard let nextContent = grid[nextPosition] else {
                    return nil
                }
                
                switch nextContent {
                case .obstruction:
                    currentDirection = currentDirection.rotate(.CW90)
                    return Step(position: currentPosition, direction: currentDirection, content: nextContent)
                default:
                    currentPosition = nextPosition
                    return Step(position: currentPosition, direction: currentDirection, content: nextContent)
                }
            }
        }
    }
}
