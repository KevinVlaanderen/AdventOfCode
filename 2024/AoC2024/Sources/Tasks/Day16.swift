import Foundation
import Framework
internal import SwiftGraph
internal import Algorithms

public final class Day16: Day<Task, Int> {
    public override func perform() throws -> Int {
        let grid = try parse(data)
        let maze = try Maze(from: grid)
        
        let reindeer = Step.move(to: maze.start, direction: .right)

        let (cameFrom, costSoFar) = try maze.path(from: reindeer, to: maze.end)
        
        let costs = costSoFar.filter { stepCost in
            if case .move(let to, _) = stepCost.key {
                return to == maze.end
            }
            return false
        }.map({ $0.value })
        let minimumCost = costs.dropFirst().reduce(costs.first!, min)
        
        switch param {
        case .task1:
            return minimumCost
        case .task2:
            var found = Set<Point>()
            
            let validSteps = cameFrom.keys.filter {
                if case .move(let to, _) = $0 {
                    return to == maze.end
                }
                return false
            }
            
            for validStep in validSteps {
                getValidPoints(for: validStep, cameFrom: cameFrom, found: &found)
            }
            
            return found.count
        }
    }
    
    private func getValidPoints(for step: Step, cameFrom: [Step: Set<Step>], found: inout Set<Point>) {
        if case .move(let to, _) = step {
            found.insert(to)
        }
        
        guard let previousSteps = cameFrom[step] else {
            return
        }
        
        for previousStep in previousSteps {
            getValidPoints(for: previousStep, cameFrom: cameFrom, found: &found)
        }
    }

    private func parse(_ data: String) throws -> any Grid<TileType> {
        try ArrayGrid(data.split(whereSeparator: \.isNewline).map { line throws in
            try Array(line).map { character throws in
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
    
    private enum TileType: Equatable, Hashable {
        case empty, wall, start, end
    }
    
    private enum Step: Equatable, Hashable {
        case move(to: Point, direction: Direction)
        case rotate(direction: Direction, at: Point)
    }
    
    private struct QueuedStep: Equatable, Comparable, Hashable {
        let step: Step
        let priority: Int
        
        static func < (lhs: QueuedStep, rhs: QueuedStep) -> Bool {
            lhs.priority < rhs.priority
        }
    }
    
    private struct Maze {
        let start: Point
        let end: Point
        let grid: any Grid<TileType>
        
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
        
        func path(from start: Step, to end: Point) throws -> ([Step: Set<Step>], [Step: Int]) {
            guard case .move(_, _) = start else {
                throw Day16Error.invalidMove
            }
            
            var frontier = PriorityQueue<QueuedStep>(ascending: true, startingValues: [QueuedStep(step: start, priority: 0)])
            
            var cameFrom: [Step: Set<Step>] = [start: []]
            var costSoFar: [Step: Int] = [start: 0]
            
            while !frontier.isEmpty {
                guard let current = frontier.pop() else {
                    break
                }
                
                let (currentPoint, currentDirection) = switch current.step {
                case let .move(to, direction): (to, direction)
                case let .rotate(direction, at): (at, direction)
                }
                
                if currentPoint == end {
                    break
                }
                
                for direction in Direction.allCases {
                    let rotate = currentDirection != direction
                    
                    let nextPoint = currentPoint.neighbour(direction: direction)
                    
                    if let nextTile = grid[nextPoint], nextTile == .wall {
                        continue
                    }
                    
                    guard let currentCost = costSoFar[current.step] else {
                        continue
                    }
                    
                    let newCost = currentCost + (rotate ? 1000 : 1)
                    
                    let nextStep: Step = rotate ? .rotate(direction: direction, at: currentPoint) : .move(to: nextPoint, direction: direction)
                    
                    guard let nextCost = costSoFar[nextStep] else {
                        costSoFar[nextStep] = newCost
                        frontier.push(QueuedStep(step: nextStep, priority: newCost))
                        cameFrom[nextStep] = [current.step]
                        continue
                    }
                    
                    if newCost > nextCost {
                        continue
                    } else if newCost < nextCost {
                        costSoFar[nextStep] = newCost
                        frontier.push(QueuedStep(step: nextStep, priority: newCost))
                        cameFrom[nextStep] = [current.step]
                    } else {
                        cameFrom[nextStep]?.insert(current.step)
                    }
                }
            }
            
            return (cameFrom, costSoFar)
        }
    }
    
    private enum Day16Error: Error {
        case invalidMove
    }
}
