import Foundation
import Framework
internal import Algorithms

public struct Day13: Day {
    public init() {}
        
    public func perform(task: Task, data: String, param: P) throws -> Int {
        let machines = parse(data)
        
        return machines.map { machine in
            var limit = min(machine.prize.x/machine.b.x, machine.prize.y/machine.b.y)
            
            while limit >= 0 {
                let remainderX = machine.prize.x-limit*machine.b.x
                let remainderY = machine.prize.y-limit*machine.b.y
                
                let fitsX = remainderX / machine.a.x
                let fitsY = remainderY / machine.a.y
                
                if remainderX % machine.a.x == 0 && remainderY % machine.a.y == 0 && fitsX == fitsY {
                    return limit+fitsX*3
                }
                
                limit -= 1
            }
            return 0
        }
        .reduce(0, +)
    }
    
    nonisolated(unsafe)
    private static let emptyLinePattern = /\n\n/
    nonisolated(unsafe)
    private static let buttonPattern = /Button .: X\+(\d+), Y\+(\d+)/
    nonisolated(unsafe)
    private static let prizePattern = /Prize: X=(\d+), Y=(\d+)/
    
    private func parse(_ data: String) -> [Machine] {
        data.split(separator: Day13.emptyLinePattern).map { block in
            let lines = block.split(whereSeparator: \.isNewline)
            guard let aMatch = lines[0].wholeMatch(of: Day13.buttonPattern) else {
                fatalError("Button A is not a match")
            }
            guard let bMatch = lines[1].wholeMatch(of: Day13.buttonPattern) else {
                fatalError("Button B is not a match")
            }
            guard let prizeMatch = lines[2].wholeMatch(of: Day13.prizePattern) else {
                fatalError("Prize is not a match")
            }
            return Machine(
                a: Button(x: Int(aMatch.output.1)!, y: Int(aMatch.output.2)!),
                b: Button(x: Int(bMatch.output.1)!, y: Int(bMatch.output.2)!),
                prize: Prize(x: Int(prizeMatch.output.1)!, y: Int(prizeMatch.output.2)!)
            )
        }
    }
    
    private struct Machine {
        let a: Button
        let b: Button
        let prize: Point
    }
    
    private typealias Button = Point
    private typealias Prize = Point
}
