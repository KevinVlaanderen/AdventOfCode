import Foundation
internal import Algorithms
internal import HeapModule
import Framework

public struct Day18: Day {
    public typealias P = (Task, Int, Int) // task, grid size, time
    public typealias R = String
    
    private let data: String
    private let param: P
    
    public init(data: String, param: P) {
        self.data = data
        self.param = param
    }
    
    public func perform() throws -> R {
        let (grid, bytes) = try parse(data: data, size: param.1)

        let start = Point(x: 0, y: 0)
        let goal = Point(x: grid.width-1, y: grid.height-1)
        
        switch param.0 {
        case .task1:
            let (cameFrom, _) = findPath(grid: grid, start: start, goal: goal, startTime: param.2)
            let path = reconstructPath(cameFrom: cameFrom, start: start, goal: goal)

            return "\(path.count)"
        case .task2:
            var lower = param.2
            var upper = bytes.count
            var mid = lower + (upper-lower)/2
            var lowestFailed: Int?
            
            while true {
                let (newCameFrom, _) = findPath(grid: grid, start: start, goal: goal, startTime: mid)
                let newPath = reconstructPath(cameFrom: newCameFrom, start: start, goal: goal)
               
                if newPath.count == 0 {
                    lowestFailed = mid
                    upper = mid
                    mid = lower + (upper-lower)/2
                } else {
                    if let lowestFailed = lowestFailed, mid+1 == lowestFailed {
                        break
                    }
                    lower = mid
                    mid = lower + (upper-lower)/2
                }
            }
            
            guard let lowestFailed = lowestFailed else {
                throw AoCError.invalidTask("no result found")
            }
            
            let point = bytes[lowestFailed-1]
            
            return "\(point.x),\(point.y)"
        }
    }
    
    nonisolated(unsafe)
    private static let bytePattern = /(\d+),(\d+)/
    
    private func parse(data: String, size: Int) throws -> (any Grid<Int>, [Point]) {
        let bytes = data.matches(of: Day18.bytePattern)
            .map({ Point(x: Int($0.output.1)!, y: Int($0.output.2)!) })
        
        var grid = ArrayGrid(width: size, height: size, defaultValue: 0)
        
        for byte in bytes.enumerated() {
            grid[byte.element] = byte.offset + 1
        }
        
        return (grid, bytes)
    }
    
    private struct PriorityPoint: Comparable {
        let point: Point
        let priority: Int
        
        static func < (lhs: Day18.PriorityPoint, rhs: Day18.PriorityPoint) -> Bool {
            lhs.priority < rhs.priority
        }
    }
    
    private func findPath(grid: any Grid<Int>, start: Point, goal: Point, startTime: Int) -> ([Point: Point?], [Point: Int]) {
        var frontier = Heap<PriorityPoint>([PriorityPoint(point: start, priority: 0)])
        
        var cameFrom: [Point: Point?] = [start: nil]
        var costSoFar: [Point: Int] = [start: startTime]
        
        while !frontier.isEmpty {
            guard let current = frontier.popMin(),
                  let currentCost = costSoFar[current.point],
                  current.point != goal else {
                break
            }
            
            let validPoints = current.point.neighbours(directions: Direction.orthogonal).filter {
                guard let time = grid[$0] else {
                    return false
                }
                return time == 0 || time > startTime
            }

            for next in validPoints {
                let newCost = currentCost + 1

                if let nextCost = costSoFar[next], newCost >= nextCost {
                    continue
                }
                
                costSoFar[next] = newCost
                frontier.insert(PriorityPoint(point: next, priority: newCost))
                cameFrom[next] = current.point
            }
        }
        
        return (cameFrom, costSoFar)
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
}
