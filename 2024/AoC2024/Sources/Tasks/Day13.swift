import Foundation
import Framework
internal import Algorithms

public final class Day13: Day<(Int, Int), Int> {
    public override func perform() throws -> Int {
        parse(data, modifier: param.0)
            .map { machine in
                let b = (machine.a.x * machine.prize.y - machine.a.y * machine.prize.x) / (machine.a.x * machine.b.y - machine.a.y * machine.b.x)
                let a = (machine.prize.x - machine.b.x * b) / machine.a.x
                
                let prize = Prize(x: a*machine.a.x + b * machine.b.x, y: a * machine.a.y + b*machine.b.y)
                
                if (param.1 >= 0 && (a > param.1 || b > param.1)) || prize != machine.prize {
                    return 0
                }
                
                return 3*a+b
            }
            .reduce(0, +)
    }
    
    nonisolated(unsafe)
    private static let emptyLinePattern = /\n\n/
    nonisolated(unsafe)
    private static let buttonPattern = /Button .: X\+(\d+), Y\+(\d+)/
    nonisolated(unsafe)
    private static let prizePattern = /Prize: X=(\d+), Y=(\d+)/
    
    private func parse(_ data: String, modifier: Int) -> [Machine] {
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
                prize: Prize(x: Int(prizeMatch.output.1)!+modifier, y: Int(prizeMatch.output.2)!+modifier)
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
