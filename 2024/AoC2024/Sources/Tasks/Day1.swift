import Foundation
internal import Algorithms
import Framework

public struct Day1: Day {
    private let data: String
    private let param: P
    
    public init(data: String, param: P) {
        self.data = data
        self.param = param
    }
    
    public func perform() throws -> R {
        let (left, right) = try parse(data: data)
        
        return switch param {
        case .task1:
            task1(left: left, right: right)
        case .task2:
            task2(left: left, right: right)
        }
    }
    
    private func parse(data: String) throws -> ([Int], [Int]) {
        let numbers = try data.split(whereSeparator: \.isWhitespace).map(toInt)

        let left = numbers.striding(by: 2)
        let right = numbers.dropFirst().striding(by: 2)
        
        return (Array(left), Array(right))
    }
    
    private func task1(left: [Int], right: [Int]) -> Int {
        zip(left.sorted(), right.sorted())
            .map({ abs($0.0 - $0.1) })
            .reduce(0, +)
    }
    
    private func task2(left: [Int], right: [Int]) -> Int {
        left.map { leftValue in
            leftValue * right.count(where: { rightValue in leftValue == rightValue })
        }.reduce(0, +)
    }
}
