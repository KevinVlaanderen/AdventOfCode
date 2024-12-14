import Foundation
import Framework
internal import Algorithms

public struct Day14: Day {
    public typealias P = (Int)

    public init() {}
        
    public func perform(task: Task, data: String, param: P) throws -> Int {
        let robots = parse(data)
        
        let seconds = param
        let width = robots.map({ $0.position.x }).sorted().last! + 1
        let height = robots.map({ $0.position.y }).sorted().last! + 1
        let (midX, midY) = (width/2, height/2)
        
        let counts = robots.reduce(into: [
            Quadrants.topLeft: 0,
            Quadrants.topRight: 0,
            Quadrants.bottomLeft: 0,
            Quadrants.bottomRight: 0,
        ]) { counts, robot in
            var x = (robot.position.x+robot.velocity.x*seconds) % width
            var y = (robot.position.y+robot.velocity.y*seconds) % height
            
            x = x<0 ? width+x : x
            y = y<0 ? height+y : y
            
            switch (x, y) {
            case (x, y) where x<midX && y<midY:
                counts[.topLeft]? += 1
            case (x, y) where x>midX && y<midY:
                counts[.topRight]? += 1
            case (x, y) where x<midX && y>midY:
                counts[.bottomLeft]? += 1
            case (x, y) where x>midX && y>midY:
                counts[.bottomRight]? += 1
            default:
                return
            }
        }.values
            
        return counts.dropFirst().reduce(counts.first!, *)
    }
    
    nonisolated(unsafe)
    private static let robotPattern = /p=(\d+),(\d+) v=(-?\d+),(-?\d+)/
    
    private func parse(_ data: String) -> [Robot] {
        data.matches(of: Day14.robotPattern).map { match in
            Robot(
                position: Position(x: Int(match.output.1)!, y: Int(match.output.2)!),
                velocity: Position(x: Int(match.output.3)!, y: Int(match.output.4)!)
            )
        }
    }
    
    private struct Robot {
        let position: Point
        let velocity: Point
    }
    
    private typealias Position = Point
    private typealias Velocity = Point
    
    private enum Quadrants {
        case topLeft, topRight, bottomLeft, bottomRight
    }
}
