import Foundation
internal import Algorithms
internal import HeapModule
import Framework

public struct Day20: Day {
    private let data: String
    private let param: P
    
    public init(data: String, param: P) {
        self.data = data
        self.param = param
    }
    
    public func perform() throws -> R {
        let grid = try parse(data: data)
        let track = try RaceTrack(from: grid)
        
        let (cameFrom, costSoFar) = try track.path(from: track.start, to: track.end)
        let path = reconstructPath(cameFrom: cameFrom, start: track.start, goal: track.end)
        
        let baseLength = path.count
        
        var cheats: [Point: Int] = [:]
        
        for step in path {
            for neighbour in step.neighbours(directions: Direction.orthogonal) {
                if cheats[neighbour] != nil || track.grid[neighbour] != .wall {
                    continue
                }
                
                var newTrack = track
                newTrack.grid[neighbour] = .empty
                
                let (cameFrom, costSoFar) = try newTrack.path(from: newTrack.start, to: newTrack.end)
                let path = reconstructPath(cameFrom: cameFrom, start: newTrack.start, goal: newTrack.end)
                
                let saved = baseLength - path.count
                if saved <= 0 {
                    continue
                }
                cheats[neighbour] = saved
            }
        }
        
        return cheats.filter({ $0.value >= 100 }).count
    }
    
    private func parse(data: String) throws -> any Grid<TileType> {
        try ArrayGrid(data.lines.map { line throws in
            try line.map { character throws in
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
    
    private struct RaceTrack {
        let start: Point
        let end: Point
        var grid: any Grid<TileType>
        
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
        
        func path(from start: Point, to end: Point) throws -> ([Point: Point?], [Point: Int]) {
            var frontier = Heap<QueuedStep>([QueuedStep(point: start, cost: 0)])
            
            var cameFrom: [Point: Point?] = [start: nil]
            var costSoFar: [Point: Int] = [start: 0]
            
            while !frontier.isEmpty {
                guard let current = frontier.popMin() else {
                    break
                }
                
                if current.point == end {
                    break
                }
                
                for direction in Direction.orthogonal {
                    let nextPoint = current.point.neighbour(direction: direction)
                    
                    guard let nextTile = grid[nextPoint], let currentCost = costSoFar[current.point] else {
                        continue
                    }
                    
                    if nextTile == .wall {
                        continue
                    }
                    
                    let newCost = currentCost+1
                    let nextCost = costSoFar[nextPoint]
                    
                    if nextCost == nil || newCost < nextCost! {
                        costSoFar[nextPoint] = newCost
                        frontier.insert(QueuedStep(point: nextPoint, cost: newCost))
                        cameFrom[nextPoint] = current.point
                        continue
                    }
                }
            }
            
            return (cameFrom, costSoFar)
        }
    }
    
    private func reconstructPath(cameFrom: [Point: Point?], start: Point, goal: Point) -> [Point] {
        var current: Point = goal
        var path: [Point] = []
        
        if cameFrom[goal] == nil {
            return path
        }
        
        while current != start {
            path.append(current)
            guard let previous = cameFrom[current], let previous = previous else {
                break
            }
            current = previous
        }
        
        return path.reversed()
    }
    
    private enum TileType: Equatable, Hashable {
        case empty, wall, start, end
    }
    
    private struct QueuedStep: Equatable, Comparable, Hashable {
        let point: Point
        let cost: Int
        
        static func < (lhs: QueuedStep, rhs: QueuedStep) -> Bool {
            lhs.cost < rhs.cost
        }
    }
    
//    private struct Cost: Comparable, Hashable {
//        let steps: Int
//        let cheats: Int
//        
//        static func < (lhs: Day20.Cost, rhs: Day20.Cost) -> Bool {
//            lhs.steps < rhs.steps ? true : lhs.steps > rhs.steps ? false : lhs.cheats < rhs.cheats
//        }
//    }
}
