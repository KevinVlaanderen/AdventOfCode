import Foundation
internal import Algorithms
internal import SwiftGraph

struct Day7: Day {
    typealias R = Int

    func task1(data: String, param: P) throws -> R {
        let equations = parse(data)

        return equations
            .filter(canSatisfy(operators: Operators.allCases))
            .map { $0.result }
            .reduce(0, +)
    }
    
    func task2(data: String, param: P) async throws -> R {
        return 0
    }
    
    func canSatisfy(operators: [Operators]) -> (Equation) -> Bool {
        return { equation in
            calculate(total: equation.numbers.first!, target: equation.result, numbers: Array(equation.numbers.dropFirst()), operators: operators)
        }
    }
    
    func calculate(total: Int, target: Int, numbers: [Int], operators: [Operators]) -> Bool {
        if numbers.count == 0 {
            return false
        }
        for op in operators {
            let newTotal = op.apply(a: total, b: numbers.first!)
            if newTotal == target {
                return true
            } else if newTotal > target {
                continue
            } else {
                if calculate(total: newTotal, target: target, numbers: Array(numbers.dropFirst()), operators: operators) {
                    return true
                }
            }
        }
        return false
    }
    
    private static let linePattern = /(\d+):((?: \d+)+)/
    
    private func parse(_ data: String) -> [Equation] {
        return data.matches(of: Day7.linePattern).map { match in
            Equation(
                result: Int(match.output.1)!,
                numbers: match.output.2.split(whereSeparator: \.isWhitespace).map({
                    Int($0.trimmingCharacters(in: .whitespaces))!
                })
            )
        }
    }
    
    struct Equation {
        let result: Int
        let numbers: [Int]
    }
    
    enum Operators: CaseIterable {
        case add, multiply
    }
}

extension Day7.Operators {
    func apply(a: Int, b: Int) -> Int {
        switch self {
        case .add:
            return a+b
        case .multiply:
            return a*b
        }
    }
}
