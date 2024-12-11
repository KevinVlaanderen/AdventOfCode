import Foundation
import Framework
internal import Algorithms

public struct Day11: Day {
    public typealias P = Int
    
    public init() {}
        
    public func perform(task: Task, data: String, param: P) throws -> Int {
        let stones = parse(data)
        
        var cache: Cache = [:]
        
        return stones.flatMap { expandStone(stone: $0, count: param, cache: &cache) }.count
    }
    
    private func expandStone(stone: Int, count: Int, cache: inout Cache) -> [Int] {
        let key = Key(stone: stone, count: count)
                      
        if count == 0 {
            return [stone]
        } else if let cachedStones = cache[key] {
            return cachedStones
        }
        
        var newStones: [Int] = []
        let numDigits = stone.usefulDigits
        
        if stone == 0 {
            newStones = expandStone(stone: 1, count: count-1, cache: &cache)
        } else if numDigits > 0 && numDigits % 2 == 0 {
            let divisor = Int(pow(10, Double(numDigits/2)))
            newStones = [
                expandStone(stone: stone / divisor, count: count-1, cache: &cache),
                expandStone(stone: stone % divisor, count: count-1, cache: &cache),
            ].flatMap { $0 }
        } else {
            newStones = expandStone(stone: stone * 2024, count: count-1, cache: &cache)
        }
        
        cache[key] = newStones

        return newStones
    }
    
    private func parse(_ data: String) -> [Int] {
        data.split(whereSeparator: \.isWhitespace).map { Int($0)! }
    }
    
    private typealias Cache = [Key: [Int]]
    
    private struct Key: Hashable {
        let stone: Int
        let count: Int
    }
}

