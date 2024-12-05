import Foundation
internal import Algorithms

struct Day4: Day {
    typealias P = String
    typealias R = Int

    func task1(data: String, param: P) throws -> R {
        let grid = parse(data)
        var count = 0
        
        for position in grid.filter(param.first!) {
            for direction in Direction.allCases {
                if wordFound(grid: grid, word: param, position: position, direction: direction) {
                    count += 1
                }
            }
        }

        return count
    }
    
    func task2(data: String, param: P) throws -> R {
        let grid = parse(data)
        
        var crossed: [Point: Int] = [:]
        
        for position in grid.filter(param.first!) {
            for direction in [Direction.NE, Direction.SE, Direction.SW, Direction.NW] {
                if wordFound(grid: grid, word: param, position: position, direction: direction) {
                    let center = position.neighbour(direction: direction)
                    let current = crossed[center] ?? 0
                    crossed[center] = current+1
                }
            }
        }

        return crossed.count(where: { $0.value == 2 })
    }
    
    private func parse(_ data: String) -> ArrayGrid<Character> {
        return ArrayGrid(data.split(whereSeparator: \.isNewline).map({ Array($0) }))
    }
    
    private func wordFound(grid: ArrayGrid<Character>, word: String, position: Point, direction: Direction) -> Bool {
        guard let currentCharacter = word.first else {
            return true
        }
        guard let gridCharacter = grid.get(position) else {
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
