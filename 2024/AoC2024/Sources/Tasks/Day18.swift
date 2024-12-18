import Foundation
internal import Algorithms
internal import SwiftGraph
import Framework

public struct Day18: Day {
    public typealias P = (Int, Int) // grid size, time
    
    private let data: String
    private let param: P
    
    public init(data: String, param: P) {
        self.data = data
        self.param = param
    }
    
    public func perform() throws -> R {
        let grid = try parse(data: data, size: param.0)

        let start = Point(x: 0, y: 0)
        let goal = Point(x: grid.width-1, y: grid.height-1)
        
        let (cameFrom, _) = findPath(grid: grid, start: start, goal: goal, startTime: param.1)

        let path = reconstructPath(cameFrom: cameFrom, start: start, goal: goal)
        
//        draw(grid: grid, path: path, upToTime: param.1)
        
        return path.count
    }
    
    nonisolated(unsafe)
    private static let bytePattern = /(\d+),(\d+)/
    
    private func parse(data: String, size: Int) throws -> any Grid<Int> {
        let bytes = data.matches(of: Day18.bytePattern)
            .map({ Point(x: Int($0.output.1)!, y: Int($0.output.2)!) })
        
        var grid = ArrayGrid(width: size, height: size, defaultValue: 0)
        
        for byte in bytes.enumerated() {
            grid[byte.element] = byte.offset + 1
        }
        
        return grid
    }
    
    private struct PriorityPoint: Comparable {
        let point: Point
        let priority: Int
        
        static func < (lhs: Day18.PriorityPoint, rhs: Day18.PriorityPoint) -> Bool {
            lhs.priority < rhs.priority
        }
    }
    
    private func findPath(grid: any Grid<Int>, start: Point, goal: Point, startTime: Int) -> ([Point: Point?], [Point: Int]) {
        var frontier = PriorityQueue<PriorityPoint>(ascending: true, startingValues: [PriorityPoint(point: start, priority: 0)])
        
        var cameFrom: [Point: Point?] = [start: nil]
        var costSoFar: [Point: Int] = [start: 0]
        
        while !frontier.isEmpty {
            guard let current = frontier.pop(),
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
                frontier.push(PriorityPoint(point: next, priority: newCost))
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
    
//    func draw(grid: any Grid<Int>, path: [Point], upToTime: Int) {
//        print(String(repeating: "-", count: grid.width))
//        for y in 0..<grid.height {
//            let line = (0..<grid.width).map {
//                let point = Point(x: $0, y: y)
//                let time = grid[point]!
//                
//                if path.contains(point) {
//                    return "x"
//                }else if time == 0 || time > upToTime {
//                    return "."
//                } else {
//                    return "#"
//                }
//            }.joined()
//            print("|" + line + "|")
//        }
//        print(String(repeating: "-", count: grid.width))
//    }
}
