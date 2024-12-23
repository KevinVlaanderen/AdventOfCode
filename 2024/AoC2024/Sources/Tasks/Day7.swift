import Foundation
internal import Algorithms
import Framework

public final class Day7: Day<[Day7.OperatorParam], Int> {
    public enum OperatorParam: CaseIterable {
        case add, multiply, combine
    }
    
    public override func perform(data: String, task: Task, param: P) throws -> Int {
        let (equations, operators) = parse(data, param: param)
        
        return equations.filter(canSatisfy(operators: operators))
            .map { $0.result }
            .reduce(0, +)
    }
    
    private func parse(_ data: String, param: P) -> ([Equation], [Operator]) {
        let linePattern = /(\d+):((?: \d+)+)/
        
        let equations = data.matches(of: linePattern).map { match in
            Equation(
                result: Int(match.output.1)!,
                numbers: match.output.2.split(whereSeparator: \.isWhitespace).map({
                    Int($0.trimmingCharacters(in: .whitespaces))!
                })
            )
        }
        
        let operators: [Operator] = param.map {
            switch $0 {
            case .add: .add
            case .combine: .combine
            case .multiply: .multiply
            }
        }
        
        return (equations, operators)
    }
    
    private func canSatisfy(operators: [Operator]) -> (Equation) -> Bool {
        { equation in
            self.calculate(total: equation.numbers.first!, target: equation.result, index: 1, numbers: equation.numbers, operators: operators)
        }
    }
    
    private func calculate(total: Int, target: Int, index: Int, numbers: [Int], operators: [Operator]) -> Bool {
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
}

private struct Equation {
    let result: Int
    let numbers: [Int]
}

private enum Operator: CaseIterable {
    case add, multiply, combine
    
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
