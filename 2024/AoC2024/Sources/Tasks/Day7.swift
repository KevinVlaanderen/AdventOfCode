import Foundation
internal import Algorithms
import Framework

public final class Day7: Day<[Day7.Operators], Int> {
    public override func perform() throws -> Int {
       parse(data)
            .filter(canSatisfy(operators: param))
            .map { $0.result }
            .reduce(0, +)
    }
    
    nonisolated(unsafe)
    private static let linePattern = /(\d+):((?: \d+)+)/
    
    private func parse(_ data: String) -> [Equation] {
        data.matches(of: Day7.linePattern).map { match in
            Equation(
                result: Int(match.output.1)!,
                numbers: match.output.2.split(whereSeparator: \.isWhitespace).map({
                    Int($0.trimmingCharacters(in: .whitespaces))!
                })
            )
        }
    }
    
    private func canSatisfy(operators: [Operators]) -> (Equation) -> Bool {
        { equation in
            self.calculate(total: equation.numbers.first!, target: equation.result, index: 1, numbers: equation.numbers, operators: operators)
        }
    }
    
    private func calculate(total: Int, target: Int, index: Int, numbers: [Int], operators: [Operators]) -> Bool {
        if total > target {
            return false
        } else if index == numbers.endIndex {
            return total == target
        }
        let next = numbers[index]
        return operators.contains { op in
            let newTotal = op.apply(a: total, b: next)
            return calculate(total: newTotal, target: target, index: index+1, numbers: numbers, operators: operators)
        }
    }
    
    public enum Operators: CaseIterable {
        case add, multiply, combine
    }
}

private struct Equation {
    let result: Int
    let numbers: [Int]
}

extension Day7.Operators {
    func apply(a: Int, b: Int) -> Int {
        switch self {
        case .add:
            return a+b
        case .multiply:
            return a*b
        case .combine:
            return a*Int(pow(10, Double(b.usefulDigits)))+b
        }
    }
}
