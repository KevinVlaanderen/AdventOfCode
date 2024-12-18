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
        let (left, right) = parse(data: data)
        
        return switch param {
        case .task1:
            task1(left: left, right: right)
        case .task2:
            task2(left: left, right: right)
        }
    }
    
    private func parse(data: String) -> ([Int], [Int]) {
        let lines = data.split(whereSeparator: \.isWhitespace).compactMap({ Int($0) })
        let (leftElements, rightElements) = lines.enumerated().partitioned(by: { $0.offset % 2 == 0 })
        
        return (leftElements.map({ $0.element }), rightElements.map({ $0.element }))
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
