import Foundation
internal import Algorithms
import Framework

public struct Day4: Day {
    public typealias P = String
    public typealias R = Int
    
    public init() {}
    
    public func perform(task: Task, data: String, param: P) throws -> Int {
        let grid = parse(data)

        switch task {
        case .task1:
            return task1(grid: grid, word: param)
        case .task2:
            precondition(param.count % 2 == 1)
            return task2(grid: grid, word: param)
        @unknown default:
            fatalError("unknown task")
        }
    }
    
    private func parse(_ data: String) -> some Grid<Character> {
        return ArrayGrid(data.split(whereSeparator: \.isNewline).map({ Array($0) }))
    }
    
    private func task1(grid: any Grid<Character>, word: String) -> R {
        var count = 0
        
        for item in grid.filter({ $0.value == word.first! }) {
            for direction in Direction.allCases {
                if wordFound(grid: grid, word: word, position: item.position, direction: direction) {
                    count += 1
                }
            }
        }

        return count
    }
    
    private func task2(grid: any Grid<Character>, word: String) -> R {
        var crossed: [Point: Int] = [:]
        
        for item in grid.filter({ $0.value == word.first! }) {
            for direction in [Direction.NE, Direction.SE, Direction.SW, Direction.NW] {
                if wordFound(grid: grid, word: word, position: item.position, direction: direction) {
                    let center = item.position.neighbour(direction: direction, distance: word.count / 2)
                    let current = crossed[center] ?? 0
                    crossed[center] = current+1
                }
            }
        }

        return crossed.count(where: { $0.value >= 2 })
    }
    
    private func wordFound(grid: any Grid<Character>, word: String, position: Point, direction: Direction) -> Bool {
        guard let currentCharacter = word.first else {
            return true
        }
        guard let gridCharacter = grid[position] else {
            return false
        }
        
        if currentCharacter != gridCharacter {
            return false
        }
        
        let newWord = String(word.dropFirst())
        let newPosition = position.neighbour(direction: direction)
        
        return wordFound(grid: grid, word: newWord, position: newPosition, direction: direction)
    }
}
