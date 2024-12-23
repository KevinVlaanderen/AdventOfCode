import Foundation
internal import Algorithms
internal import SwiftGraph
import Framework

public final class Day6: Day<Void, Int> {
    public override func perform(data: String, task: Task, param: P) throws -> Int {
        let grid: any Grid<Content> = parse(data)

        guard let (guardPosition, guardDirection) = findGuard(grid: grid) else {
            fatalError("no guard found")
        }
        
        return switch task {
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
                case "^":   .patrolGuard(.north)
                default:    fatalError("unknown grid item: \(character)")
                }
            }
        })
    }
    
    private func findGuard(grid: any Grid<Content>) -> (Point, Heading)? {
        for item in grid {
            if case let .patrolGuard(guardDirection) = item.value {
                return (item.position, guardDirection)
            }
        }
        return nil
    }
    
    private func task1(grid: any Grid<Content>, startPosition: Point, startDirection: Heading) throws -> Int {
        PathTracker(grid: grid, startPosition: startPosition, startDirection: startDirection)
            .map { step in step.position }
            .uniqued()
            .count { _ in true }
    }
    
    private func task2(grid: any Grid<Content>, startPosition: Point, startDirection: Heading) throws -> Int {
        PathTracker(grid: grid, startPosition: startPosition, startDirection: startDirection)
            .filter(isValidTargetForObstruction(grid: grid, guardPosition: startPosition))
            .uniqued(on: nextPosition)
            .count(where: entersLoop(grid: grid))
    }
    
    private func isValidTargetForObstruction(grid: any Grid<Content>, guardPosition: Point) -> (Step) -> Bool {
        { step in
            let nextPosition = self.nextPosition(for: step)
            guard let content = grid[nextPosition], content != .obstruction else {
                return false
            }
            return nextPosition != guardPosition
        }
    }
    
    private func nextPosition(for step: Step) -> Point {
        step.position.neighbour(heading: step.heading)
    }
    
    private func entersLoop(grid: any Grid<Content>) -> (Step) -> Bool {
        { outerStep in
            var newGrid = grid
            newGrid[outerStep.position.neighbour(heading: outerStep.heading)] = .obstruction
            
            var seen: [Step] = []
            
            for (step, nextStep) in PathTracker(grid: newGrid, startPosition: outerStep.position, startDirection: outerStep.heading).adjacentPairs() {
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
}

private enum Content: Equatable, Hashable {
    case empty, obstruction
    case patrolGuard(Heading)
}

private struct Step: Hashable {
    let position: Point
    let heading: Heading
    let content: Content
}

private struct PathTracker: Sequence {
    let grid: any Grid<Content>
    let startPosition: Point
    let startDirection: Heading
    
    func makeIterator() -> PathIterator {
        PathIterator(grid: grid, currentPosition: startPosition, currentHeading: startDirection)
    }
    
    struct PathIterator: IteratorProtocol {
        let grid: any Grid<Content>
        var currentPosition: Point
        var currentHeading: Heading
        var started = false
        
        mutating func next() -> Step? {
            if !started {
                guard let currentContent = grid[currentPosition] else {
                    return nil
                }
                started = true
                return Step(position: currentPosition, heading: currentHeading, content: currentContent)
            }

            let nextPosition = currentPosition.neighbour(heading: currentHeading)

            guard let nextContent = grid[nextPosition] else {
                return nil
            }
            
            switch nextContent {
            case .obstruction:
                currentHeading = currentHeading.rotate(.clockwise90)
                return Step(position: currentPosition, heading: currentHeading, content: nextContent)
            default:
                currentPosition = nextPosition
                return Step(position: currentPosition, heading: currentHeading, content: nextContent)
            }
        }
    }
}
