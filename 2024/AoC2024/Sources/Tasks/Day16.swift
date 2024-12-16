import Foundation
import Framework
internal import SwiftGraph
internal import Algorithms

public struct Day16: Day {
    private let data: String
    private let param: P
    
    public init(data: String, param: P) {
        self.data = data
        self.param = param
    }
        
    public func perform() throws -> Int {
        let grid = try parse(data)
        let maze = try Maze(from: grid)
        
        let reindeer = Reindeer(point: maze.start, direction: .E)

        let (_, costSoFar) = maze.path(from: reindeer, to: maze.end)
        
        let costs = costSoFar.filter { $0.key.point == maze.end }.map({ $0.value })
        
        return costs.dropFirst().reduce(costs.first!, min)
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
    
    private struct Position: Equatable, Hashable {
        let point: Point
        let direction: Direction
    }
    
    private struct Step: Equatable, Comparable, Hashable {
        let position: Position
        let cost: Int
        
        static func < (lhs: Step, rhs: Step) -> Bool {
            lhs.cost < rhs.cost
        }
    }
    
    private typealias Reindeer = Position
    
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
        
        func path(from start: Position, to end: Point) -> ([Position: Position?], [Position: Int]) {
            var frontier = PriorityQueue<Step>(ascending: true, startingValues: [Step(position: start, cost: 0)])
            
            var cameFrom: [Position: Position?] = [start: nil]
            var costSoFar: [Position: Int] = [start: 0]
            
            while !frontier.isEmpty {
                guard let current = frontier.pop() else {
                    break
                }
                
                if current.position.point == end {
                    break
                }
                
                for next in neighbours(of: current.position, towards: Direction.orthogonal) {
                    guard let currentCost = costSoFar[current.position] else {
                        continue
                    }
                    
                    let newCost = currentCost + cost(from: current.position, to: next)
                    
                    if let nextCost = costSoFar[next], newCost >= nextCost {
                        continue
                    }
                    
                    costSoFar[next] = newCost
                    // priority
                    frontier.push(Step(position: next, cost: newCost))
                    cameFrom[next] = current.position
                }
            }
            
            return (cameFrom, costSoFar)
        }
        
        func neighbours(of position: Position, towards directions: [Direction]) -> [Position] {
            grid.neighbours(of: position.point, towards: Direction.orthogonal).compactMap { neighbour in
                if let neighbourType = grid[neighbour.position], neighbourType == .wall {
                    return nil
                }
                        
                let direction: Direction? =
                if neighbour.position.x == position.point.x && neighbour.position.y < position.point.y { Direction.N }
                else if neighbour.position.x > position.point.x && neighbour.position.y == position.point.y { Direction.E }
                else if neighbour.position.x == position.point.x && neighbour.position.y > position.point.y { Direction.S }
                else if neighbour.position.x < position.point.x && neighbour.position.y == position.point.y { Direction.W }
                else { nil }
                
                guard let direction = direction else {
                    return nil
                }
                
                return Position(point: neighbour.position, direction: direction)
            }
        }
        
        func cost(from: Position, to: Position) -> Int {
            from.direction == to.direction ? 1 : 1001
        }
    }
    
    private enum TileType: Equatable, Hashable {
        case empty, wall, start, end
    }
    
    
}
