import Foundation
internal import Algorithms
internal import SwiftGraph
import Framework

public struct Day7: Day {
    public typealias P = [Operators]
    
    public init() {}
    
    public func perform(task: Task, data: String, param: P) throws -> Int {
       return parse(data)
            .filter(canSatisfy(operators: param))
            .map { $0.result }
            .reduce(0, +)
    }
    
    func canSatisfy(operators: [Operators]) -> (Equation) -> Bool {
        return { equation in
            calculate(total: equation.numbers.first!, target: equation.result, index: 1, numbers: equation.numbers, operators: operators)
        }
    }
    
    func calculate(total: Int, target: Int, index: Int, numbers: [Int], operators: [Operators]) -> Bool {
        if index >= numbers.count {
            return total == target
        }
        for op in operators {
            let newTotal = op.apply(a: total, b: numbers[index])
            if newTotal == target {
                return true
            } else if newTotal > target {
                continue
            } else {
                if calculate(total: newTotal, target: target, index: index+1, numbers: numbers, operators: operators) {
                    return true
                }
            }
        }
        return false
    }
    
    nonisolated(unsafe)
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
    
    public enum Operators: CaseIterable, Sendable {
        case add, multiply, combine
    }
}

extension Day7.Operators {
    func apply(a: Int, b: Int) -> Int {
        switch self {
        case .add:
            return a+b
        case .multiply:
            return a*b
        case .combine:
            let total = a*Int(pow(10, Double(b.usefulDigits)))+b
            return total
        }
    }
}
