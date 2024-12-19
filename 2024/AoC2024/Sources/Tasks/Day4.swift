import Foundation
internal import Algorithms
import Framework

public struct Day4: Day {
    public typealias P = (task: Task, word: String)
    
    private let data: String
    private let param: P
    
    public init(data: String, param: P) {
        self.data = data
        self.param = param
    }
    
    public func perform() throws -> R {
        let grid = parse(data)

        return switch param.task {
        case .task1: try task1(grid: grid, word: param.word)
        case .task2: try task2(grid: grid, word: param.word)
        }
    }
    
    private func parse(_ data: String) -> some Grid<Character> {
        ArrayGrid(data.lines.map(toCharacters))
    }
    
    private func task1(grid: any Grid<Character>, word: String) throws -> Int {
        guard let firstCharacter = word.first else {
            throw AoCError.invalidState("word is empty")
        }
        
        var count = 0
        
        for item in grid {
            if item.value != firstCharacter {
                continue
            }
            
            for direction in Direction.allCases {
                if wordFound(grid: grid, word: word, position: item.position, direction: direction) {
                    count += 1
                }
            }
        }

        return count
    }
    
    private func task2(grid: any Grid<Character>, word: String) throws -> Int {
        if word.count.isEven {
            throw AoCError.invalidTask("word should have an odd length")
        }
        guard let firstCharacter = word.first else {
            throw AoCError.invalidState("word is empty")
        }
        
        var crossed: [Point: Int] = [:]
        
        for item in grid {
            if item.value != firstCharacter {
                continue
            }
            
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
