import Foundation
import Framework
internal import Algorithms

public struct Day11: Day {
    public typealias P = Int
    
    public init() {}
        
    public func perform(task: Task, data: String, param: P) throws -> Int {
        let stones = parse(data)
        
        var cache: Cache = [:]
        
        return countStones(stones: stones, iterations: param, cache: &cache)
    }
    
    private func countStones(stones: [Stone], iterations: Int, cache: inout Cache) -> Int {
        if iterations == 0 {
            return stones.count
        }
        
        return stones.reduce(0) { result, stone in
            let key = Key(stone: stone, count: iterations)

            if let cachedResult = cache[key] {
                return result + cachedResult
            }
            
            let numDigits = stone.usefulDigits
            var numStones = 0
            
            if stone == 0 {
                numStones = countStones(stones: [1], iterations: iterations-1, cache: &cache)
            } else if numDigits > 0 && numDigits % 2 == 0 {
                let divisor = Int(pow(10, Double(numDigits/2)))
                let newStones = [
                    stone / divisor,
                    stone % divisor
                ]
                numStones = countStones(stones: newStones, iterations: iterations-1, cache: &cache)
            } else {
                numStones = countStones(stones: [stone*2024], iterations: iterations-1, cache: &cache)
            }
            
            cache[key] = numStones
            return result + numStones
        }
    }
    
    private func parse(_ data: String) -> [Stone] {
        data.split(whereSeparator: \.isWhitespace).map { Stone($0)! }
    }
    
    private typealias Stone = Int
    private typealias Cache = [Key: Stone]
    
    private struct Key: Hashable {
        let stone: Stone
        let count: Int
    }
}

