import Foundation
internal import Algorithms
import Framework

public struct Day19: Day {
    private let data: String
    private let param: P
    
    public init(data: String, param: P) {
        self.data = data
        self.param = param
    }
    
    public func perform() throws -> R {
        let (towels, designs) = parse(data: data)
        
        var cache: [Design: Int] = [:]
        
        return switch param {
        case .task1:
            designs.count(where: { countOptions(design: $0, towels: towels, cache: &cache) > 0 })
        case .task2:
            designs.reduce(0) { $0 + countOptions(design: $1, towels: towels, cache: &cache) }
        }
    }
    
    private func parse(data: String) -> ([Towel], [Design]) {
        let blocks = data.split(separator: "\n\n")
        
        let towels = blocks[0].split(separator: ", ").map({ Towel($0) })
        let designs = blocks[1].split(whereSeparator: \.isNewline).map({ Design($0) })
        
        return (towels, designs)
    }
    
    private func countOptions(design: Design, towels: [Towel], cache: inout [Design: Int]) -> Int {
        if design.isEmpty {
            return 1
        } else if let cachedCount = cache[design] {
            return cachedCount
        }
        
        let count = towels.reduce(0) { result, towel in
            if !design.hasPrefix(towel) {
                return result
            }
            return result+countOptions(design: design.dropFirst(towel.count), towels: towels, cache: &cache)
        }
        cache[design] = count
        return count
    }
    
    private typealias Towel = String
    private typealias Design = Substring
}
