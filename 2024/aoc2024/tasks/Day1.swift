import Foundation
internal import Algorithms

struct Day1: Day {
    typealias R = Int
    
    func task1(data: String, param: P) throws -> R {
        let (left, right) = parse(data: data)

        return zip(left.sorted(), right.sorted()).map({ abs($0.0 - $0.1) }).reduce(0, +)
    }
    
    func task2(data: String, param: P) throws -> R {
        let (left, right) = parse(data: data)
        
        return left.map({ leftValue in leftValue * right.count(where: { rightValue in leftValue == rightValue }) }).reduce(0, +)
    }
    
    private func parse(data: String) -> ([Int], [Int]) {
        let lines = data.split(whereSeparator: \.isWhitespace).map({ Int($0)! })
        let (leftElements, rightElements) = lines.enumerated().partitioned(by: { $0.offset % 2 == 0 })
        
        return (leftElements.map({ $0.element }), rightElements.map({ $0.element }))
    }

}
