import Foundation
import Framework
internal import Algorithms

public final class Day13: Day<(modifier: Int, presses: Int), Int> {
    public override func perform(data: String, param: P) throws -> Int {
        parse(data, param: param).map { machine in
            let b = (machine.a.x * machine.prize.y - machine.a.y * machine.prize.x) / (machine.a.x * machine.b.y - machine.a.y * machine.b.x)
            let a = (machine.prize.x - machine.b.x * b) / machine.a.x
            
            let prize = Prize(x: a*machine.a.x + b * machine.b.x, y: a * machine.a.y + b*machine.b.y)
            
            if (param.presses >= 0 && (a > param.presses || b > param.presses)) || prize != machine.prize {
                return 0
            }
            
            return 3*a+b
        }.reduce(0, +)
    }
    
    private func parse(_ data: String, param: P) -> [Machine] {
        let emptyLinePattern = /\n\n/
        let buttonPattern = /Button .: X\+(\d+), Y\+(\d+)/
        let prizePattern = /Prize: X=(\d+), Y=(\d+)/
        
        return data.split(separator: emptyLinePattern).map { block in
            let lines = block.split(whereSeparator: \.isNewline)
            guard let aMatch = lines[0].wholeMatch(of: buttonPattern) else {
                fatalError("Button A is not a match")
            }
            guard let bMatch = lines[1].wholeMatch(of: buttonPattern) else {
                fatalError("Button B is not a match")
            }
            guard let prizeMatch = lines[2].wholeMatch(of: prizePattern) else {
                fatalError("Prize is not a match")
            }
            return Machine(
                a: Button(x: Int(aMatch.output.1)!, y: Int(aMatch.output.2)!),
                b: Button(x: Int(bMatch.output.1)!, y: Int(bMatch.output.2)!),
                prize: Prize(x: Int(prizeMatch.output.1)!+param.modifier, y: Int(prizeMatch.output.2)!+param.modifier)
            )
        }
    }
}

private struct Machine {
    let a: Button
    let b: Button
    let prize: Point
}

private typealias Button = Point
private typealias Prize = Point
