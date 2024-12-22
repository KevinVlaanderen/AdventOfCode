import Foundation
import Framework
internal import Algorithms

public final class Day14: Day<(task: Task, draw: Bool), Int> {
    public override func perform(data: String, param: P) throws -> Int {
        var room = Room(robots: parse(data))
        
        return switch param.task {
        case .task1: task1(room: &room, draw: param.draw)
        case .task2: task2(room: &room, draw: param.draw)
        }
    }
    
    private func task1(room: inout Room, draw: Bool) -> Int {
        room.step(amount: 100)

        if draw {
            room.draw()
        }
        
        return room.quadrants().filter({ $0.key != .none }).reduce(0) { result, quadrant in
            result == 0 ? quadrant.value.count : result * quadrant.value.count
        }
    }
    
    private func task2(room: inout Room, draw: Bool) -> Int {
        if !draw {
            return 7492
        }
        
        var step = 0
        while true {
            room.step(amount: 1)
            print(step)
            room.draw()
            step += 1
        }
    }

    private func parse(_ data: String) -> [Robot] {
        let robotPattern = /p=(\d+),(\d+) v=(-?\d+),(-?\d+)/
        
        return data.matches(of: robotPattern).map { match in
            Robot(
                position: Position(x: Int(match.output.1)!, y: Int(match.output.2)!),
                velocity: Position(x: Int(match.output.3)!, y: Int(match.output.4)!)
            )
        }
    }
}

private struct Room {
    var robots: [Robot]
    let width: Int
    let height: Int
    let midX: Int
    let midY: Int
    
    init(robots: [Robot]) {
        self.robots = robots
        self.width = robots.map({ $0.position.x }).sorted().last! + 1
        self.height = robots.map({ $0.position.y }).sorted().last! + 1
        self.midX = width/2
        self.midY = height/2
    }
    
    mutating func step(amount: Int) {
        for index in 0..<robots.count {
            var robot = robots[index]
            robot.position.x = (robot.position.x+robot.velocity.x*amount) % width
            robot.position.y = (robot.position.y+robot.velocity.y*amount) % height
            
            robot.position.x = robot.position.x<0 ? width+robot.position.x : robot.position.x
            robot.position.y = robot.position.y<0 ? height+robot.position.y : robot.position.y
            
            robots[index] = robot
        }
    }
    
    func quadrants() -> [Quadrants: [Robot]] {
        robots.grouped { robot in
            if robot.position.x<midX && robot.position.y<midY {
                return .topLeft
            } else if robot.position.x>midX && robot.position.y<midY {
                return .topRight
            } else if robot.position.x<midX && robot.position.y>midY {
                return .bottomLeft
            } else if robot.position.x>midX && robot.position.y>midY {
                return .bottomRight
            } else {
                return .none
            }
        }
    }
    
    func draw() {
        var grid = ArrayGrid(width: width, height: height, defaultValue: " ")
        for robot in robots {
            grid[Point(x: robot.position.x, y: robot.position.y)] = "x"
        }
        for y in 0..<height {
            print("|" + (0..<width).map({ grid[Point(x: $0, y: y)]!}).joined() + "|")
        }
        print(String(repeating: "-", count: width))
    }
}

private struct Robot {
    var position: Point
    var velocity: Point
}

private typealias Position = Point
private typealias Velocity = Point

private enum Quadrants {
    case topLeft, topRight, bottomLeft, bottomRight, none
}
