import Foundation
import Framework
internal import HeapModule
internal import BitCollections

public final class Day25: Day<Void, Int> {
    public override func perform(data: String, task: Task, param: P) throws -> R {
        let (locks, keys) = try parse(data)
        
        return locks.reduce(0) { result, lock in
            result + keys.count { key in
                key.enumerated().allSatisfy({ $0.element + lock[$0.offset] <= 6})
            }
        }
    }
    
    private func parse(_ data: String) throws -> ([[Int]], [[Int]]) {
        let blocks = data.blocks.map({ $0.lines.map({ $0.characters }) })
        
        let locks = blocks.filter(isType(.lock)).map(toScore)
        let keys = blocks.filter(isType(.key)).map(toScore)
        
        return (locks, keys)
    }
    
    private func isType(_ type: BlockType) -> ([[Character]]) -> Bool {
        { block in block.first?.allSatisfy({ $0 == type.rawValue }) ?? false }
    }
    
    
    private func toScore(block: [[Character]]) -> [Int] {
        var score: [Int] = []
        for x in 0..<block.first!.count {
            score.append(block.dropFirst().count(where: { $0[x] == "#" }))
        }
        return score
    }
    
    private enum BlockType: Character {
        case lock = "#"
        case key = "."
    }
}
