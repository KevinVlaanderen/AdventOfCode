import Foundation
internal import Algorithms
import Framework

public struct Day4: Day {
    public typealias P = (Task, String)
    
    private let data: String
    private let param: P
    
    public init(data: String, param: P) {
        self.data = data
        self.param = param
    }
    
    public func perform() throws -> Int {
        let grid = parse(data)

        switch param.0 {
        case .task1:
            return task1(grid: grid, word: param.1)
        case .task2:
            precondition(param.1.count % 2 == 1)
            return task2(grid: grid, word: param.1)
        }
    }
    
    private func parse(_ data: String) -> some Grid<Character> {
        return ArrayGrid(data.split(whereSeparator: \.isNewline).map({ Array($0) }))
    }
    
    private func task1(grid: any Grid<Character>, word: String) -> R {
        var count = 0
        
        for item in grid {
            if item.value != word.first! {
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
    
    private func task2(grid: any Grid<Character>, word: String) -> R {
        var crossed: [Point: Int] = [:]
        
        for item in grid {
            if item.value != word.first! {
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
